package jobscheduler

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// Job represents a schedulable task
type Job struct {
	ID       string
	Name     string
	Schedule string // Cron schedule expression
	Task     func() error
	LastRun  time.Time
	NextRun  time.Time
	Status   string // "pending", "running", "completed", "failed"
	Error    error
}

// Scheduler manages job scheduling and execution
type Scheduler struct {
	cron   *cron.Cron
	jobs   map[string]*Job
	mu     sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
	logger *log.Logger
}

// NewScheduler creates a new scheduler instance
func NewScheduler(logger *log.Logger) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		cron:   cron.New(),
		jobs:   make(map[string]*Job),
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

// AddJob schedules a new job
func (s *Scheduler) AddJob(job *Job) (cron.EntryID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate job
	if job.ID == "" || job.Schedule == "" || job.Task == nil {
		return 0, fmt.Errorf("invalid job configuration")
	}

	// Add job to cron
	id, err := s.cron.AddFunc(job.Schedule, func() {
		s.runJob(job)
	})
	if err != nil {
		return 0, fmt.Errorf("failed to schedule job: %w", err)
	}

	// Update job status
	job.Status = "pending"
	s.jobs[job.ID] = job

	// Update next run time
	if entry := s.cron.Entry(id); entry.Valid() {
		job.NextRun = entry.Next
	}

	s.logger.Printf("Added job %s with schedule %s", job.Name, job.Schedule)
	return id, nil
}

// runJob executes a job and updates its status
func (s *Scheduler) runJob(job *Job) {
	s.mu.Lock()
	job.Status = "running"
	job.LastRun = time.Now()
	s.mu.Unlock()

	s.logger.Printf("Starting job %s", job.Name)
	err := job.Task()

	s.mu.Lock()
	defer s.mu.Unlock()

	if err != nil {
		job.Status = "failed"
		job.Error = err
		s.logger.Printf("Job %s failed: %v", job.Name, err)
	} else {
		job.Status = "completed"
		job.Error = nil
		s.logger.Printf("Job %s completed successfully", job.Name)
	}

	// Update next run time
	if entry := s.cron.Entry(s.cron.Entries()[0]); entry.Valid() {
		job.NextRun = entry.Next
	}
}

// GetJobStatus returns the current status of a job
func (s *Scheduler) GetJobStatus(id string) (*Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	job, exists := s.jobs[id]
	return job, exists
}

// Start begins the scheduler
func (s *Scheduler) Start() {
	s.cron.Start()
	s.logger.Println("Scheduler started")
}

// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	s.cancel()
	s.cron.Stop()
	s.logger.Println("Scheduler stopped")
}

// Example usage
func main() {
	logger := log.New(os.Stdout, "scheduler: ", log.LstdFlags)
	scheduler := NewScheduler(logger)

	// Example job
	job := &Job{
		ID:       "job1",
		Name:     "Example Job",
		Schedule: "*/5 * * * *", // Every 5 minutes
		Task: func() error {
			fmt.Println("Executing example job")
			return nil
		},
	}

	// Add job to scheduler
	if _, err := scheduler.AddJob(job); err != nil {
		logger.Fatal(err)
	}

	// Start scheduler
	scheduler.Start()

	// Keep running until interrupted
	select {}
}
