func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }

    initColor := image[sr][sc]

    var dfs func(r int, c int)
    dfs = func (r int, c int) {
        if image[r][c] == initColor {
            image[r][c] = color
            if r >= 1 {
                dfs(r-1, c)
            }
            if r+1 < len(image) {
                dfs(r+1, c)
            }
            if c >= 1 {
                dfs(r, c-1)
            }
            if c+1 < len(image[0]) {
                dfs(r, c+1)
            }
        }
    }

    dfs(sr, sc)

    return image
}