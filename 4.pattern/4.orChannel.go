// To combine one or more done channels into a single done 2.channel that
// closes if any of its component channels close.

package main

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {  // Base case
		case 0:  // Termination criteria
			return nil
		case 1:  // Second termination criteria
			return channels[0]
		}

		orDone := make(chan interface{})

		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:  // Every recursive call to or will at least have two channels
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <- or(append(channels[3:], orDone)...):
				}
			}
		}()

		return orDone
	}
}