from collections import deque

def canEscape(grid):
    rows = len(grid)
    if rows == 0:
        return False
    cols = len(grid[0])
    
    # Find initial positions
    person = None
    fire = None
    exit_pos = None
    
    for i in range(rows):
        for j in range(cols):
            if grid[i][j] == 'P':
                person = (i, j)
            elif grid[i][j] == 'F':
                fire = (i, j)
            elif grid[i][j] == 'E':
                exit_pos = (i, j)
    
    if not person or not fire or not exit_pos:
        return False
    
    # Precompute fire arrival times
    # row = 3, col = 4
    # [
    #   [inf, inf, inf, inf],
    #   [inf, inf, inf ,inf],
    #   [inf, inf, inf, inf]
    # ]
    fire_time = [[float('inf')] * cols for _ in range(rows)] 
    fire_queue = deque()
    fire_queue.append((fire[0], fire[1]))
    fire_time[fire[0]][fire[1]] = 0
    
    dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    
    while fire_queue:
        x, y = fire_queue.popleft()
        for dx, dy in dirs:
            nx, ny = x + dx, y + dy
            if 0 <= nx < rows and 0 <= ny < cols:
                if grid[nx][ny] != 'W' and fire_time[nx][ny] == float('inf'):
                    fire_time[nx][ny] = fire_time[x][y] + 1
                    fire_queue.append((nx, ny))
    
    # BFS for person
    visited = [[False] * cols for _ in range(rows)]
    person_queue = deque()
    person_queue.append((person[0], person[1], 0))  # (x, y, time)
    visited[person[0]][person[1]] = True
    
    while person_queue:
        x, y, time = person_queue.popleft()
        
        # Check if reached exit
        if (x, y) == exit_pos:
            if time < fire_time[x][y]:
                return True
            else:
                continue  # exit is on fire at or before arrival
        
        for dx, dy in dirs:
            nx, ny = x + dx, y + dy
            if 0 <= nx < rows and 0 <= ny < cols:
                if not visited[nx][ny] and grid[nx][ny] != 'W':
                    # Person arrives at new cell at time + 1
                    if time + 1 < fire_time[nx][ny]:
                        visited[nx][ny] = True
                        person_queue.append((nx, ny, time + 1))
    
    return False