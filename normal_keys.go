package main

var (
	ignoreKeys = map[Key]bool{}
	activeKeys = map[Key]bool{}
)

func HandleNormalKeys(keys []byte) {
	pressed := make(map[Key]bool)
	ignored := make(map[Key]bool)
	for _, k := range keys {

		key := Key(k)

		switch {
		case key == 0:
			// do nothing
		case key >= KeyF1 && key <= KeyF12, key >= Key1 && key <= Key6:
			switch {
			case ignoreKeys[key] && ignored[key]:
				pressed[key] = true
				if !activeKeys[key] {
					DoKeyDown(Key(key))
				}
			case ignoreKeys[key]:
				ignored[key] = true
			default:
				pressed[key] = true
				if !activeKeys[key] {
					DoKeyDown(Key(key))
				}
			}
		default:
			pressed[key] = true
			if !activeKeys[key] {
				DoKeyDown(Key(key))
			}
		}
	}

	for key := range activeKeys {
		if !pressed[key] {
			DoKeyUp(Key(key))
		}
	}

	activeKeys = pressed
}
