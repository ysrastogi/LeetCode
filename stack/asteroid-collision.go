func asteroidCollision(asteroids []int) []int {
	st := []int{}

	for _, a := range asteroids {
		collided := false

		for len(st) > 0 && a < 0 && st[len(st)-1] > 0 {
			top := st[len(st)-1]

			if top < -a {
				st = st[:len(st)-1]
				continue
			} else if top == -a {
				// both explode
				st = st[:len(st)-1]
				collided = true
				break
			} else {
				// current asteroid explodes
				collided = true
				break
			}
		}

		if !collided {
			st = append(st, a)
		}
	}

	return st
}
