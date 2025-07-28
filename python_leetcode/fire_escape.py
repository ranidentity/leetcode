from collections import deque

def canEscape(grid):
    rows=len(grid)
    if rows == 0:
        return False
    cols = len(grid[0])

    #find initial position
    person_pos = None
    fire_pos = None
    exit_pos = None

    for i in range(rows):
        for j in range(cols):
            if grid[i][j] == 'P':
                person_pos = (i,j)
            elif grid[i][j] == 'F':
                fire_pos = (i,j)
            elif grid[i][j] == 'E':
                exit_pos = (i,j)
    if not person_pos or not fire_pos or not exit_pos:
        return False
    
    # movement direction
    dirs = [(-1,0),(1,0),(0,-1),(0,1)]

    #Queue for BFS
    person_queue = deque([(person_pos[0],person_pos[1])])
    fire_queue = deque([(fire_pos[0],fire_pos[1])])

    # To keep track of visited cells and burning cells
    visited = [[False for _ in range(cols)] for _ in range(rows)]
    fire_spread = [[False for _ in range(cols)] for _ in range(rows)]
    fire_spread[fire_pos[0]][fire_pos[1]] = True
    
    time = 0

    while person_queue:
        # Person moves first
        level_size = len(person_queue)
        for _ in range(level_size):
            x,y = person_queue.popleft()

            if fire_spread[x][y]:
                continue

            if (x,y) == exit_pos:
                return True
            
            for dx, dy in dirs:
                nx, ny = x+dx, y+dy
                if 0 <= nx < rows and 0 <= ny < cols:
                    if not visited[nx][ny] and grid[nx][ny] != 'W' and not fire_spread[nx][ny]:
                        visited[nx][ny] = True
                        person_queue.append((nx,ny))
        #then fire spread
        fire_size = len(fire_queue)
        for _ in range(fire_size):
            fx, fy = fire_queue.popleft()
            for dx,dy in dirs:
                nfx,nfy = fx+dx, fy+dy
                if 0 <= nfx < rows and 0 <= nfy<cols:
                    if not fire_spread[nfx][nfy] and grid[nfx][nfy] != 'W':
                        fire_spread[nfx][nfy] = True
                        fire_queue.append((nfx,nfy))
        
        time+= 1
    return False