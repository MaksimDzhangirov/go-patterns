package main

import (
	"fmt"
	"strings"
	"sync"
)

var data = []string{
	`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus dictum elementum nisi nec pretium. Sed eu ante tortor. Donec ex urna, volutpat sed vulputate auctor, efficitur eu sapien. Donec luctus, mi nec varius sollicitudin, magna nulla egestas massa, posuere laoreet lectus mi ut diam. Sed eget odio vitae arcu euismod dignissim. Etiam faucibus vestibulum mi vitae posuere. Aliquam non rutrum mauris, euismod ornare turpis. Curabitur vel nulla mollis, rutrum nisl nec, elementum ex. In venenatis urna at tincidunt bibendum. Aliquam mattis imperdiet sapien non tempor. In in ante condimentum, auctor eros vitae, pretium augue.`,
	//`Morbi condimentum quam non dolor pulvinar auctor. Maecenas et interdum elit, ac elementum eros. Etiam et diam porta, dignissim massa vitae, semper velit. Vestibulum eu tristique nulla. Maecenas pharetra fringilla imperdiet. Mauris lectus lorem, fermentum sed ultricies vel, consectetur ac tellus. Aliquam vehicula ante porttitor urna dignissim iaculis. Nulla ultricies arcu dolor, at aliquam arcu ullamcorper sed. Sed consectetur nisl pulvinar, dictum velit et, imperdiet diam. Nulla hendrerit risus ac laoreet accumsan. Vivamus eros diam, fringilla quis molestie sed, lobortis lacinia neque. Integer sed consequat diam. Quisque augue ipsum, imperdiet non porttitor interdum, scelerisque nec ligula. Suspendisse quis auctor ipsum, vel faucibus eros. Praesent tincidunt a risus a elementum. Phasellus pulvinar tortor ac nisl aliquam euismod.`,
	`Integer feugiat, libero a interdum tincidunt, orci massa viverra augue, eget convallis lectus purus a eros. Curabitur elementum erat justo, id ultrices velit efficitur quis. Vestibulum at scelerisque leo. Ut auctor sagittis quam congue tincidunt. Praesent malesuada mattis arcu, nec imperdiet felis fringilla sed. Aliquam viverra eleifend elementum. Quisque nec facilisis quam. Nunc nec quam sapien. Nullam congue, nisi at elementum lacinia, sem risus consectetur felis, id consectetur dui nunc non metus. Vivamus non dapibus elit. Integer eget cursus lacus, non semper elit. Cras malesuada a sapien eu euismod. Nulla eget libero lacus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean facilisis lectus ut felis viverra, eu posuere dolor hendrerit.`,
	`Integer dapibus pharetra ex a tristique. Ut pulvinar iaculis ex, id commodo ex venenatis sed. Suspendisse varius massa odio, id egestas enim aliquet vitae. Ut vitae libero massa. Phasellus nec urna ac lacus sagittis suscipit. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Proin vehicula blandit urna, sed bibendum felis dapibus quis. Sed nunc elit, ultrices vel auctor ac, ultrices sed tortor. Praesent at nibh at massa volutpat porttitor at in massa. Suspendisse non bibendum nunc.`,
	//`Quisque consectetur, ante maximus sagittis pretium, erat sapien viverra mauris, at rutrum dui ipsum at est. Mauris pharetra erat justo, ac convallis risus commodo id. Aenean interdum ultricies mi nec lobortis. Mauris pellentesque sem eu mauris facilisis luctus. Nunc tellus diam, rhoncus non gravida eget, sodales vitae nunc. Donec urna eros, finibus vel eros non, consequat vestibulum risus. Sed in orci at enim rhoncus rhoncus at vel augue. Nullam sit amet urna sed arcu ullamcorper accumsan et et tellus. Duis tincidunt consectetur ex sed laoreet. Donec facilisis elit eros, ut lobortis nisi rutrum quis. Vestibulum posuere rhoncus arcu vel efficitur. Mauris mollis a purus et convallis. Nulla facilisi. Duis quis velit non neque dapibus scelerisque sed ut nibh.`,
	`Aliquam erat volutpat. Phasellus facilisis consectetur risus, non facilisis neque gravida eget. Interdum et malesuada fames ac ante ipsum primis in faucibus. Mauris efficitur in nisl eget vestibulum. Nam egestas turpis eget est efficitur efficitur. Sed lobortis mauris at convallis tincidunt. In sed nisi ac elit feugiat gravida ut eu augue. Aliquam pretium tincidunt dolor vel accumsan. Maecenas lectus erat, tempus id fringilla non, consectetur in massa. Aliquam vel metus nec ipsum efficitur vestibulum.`,
	`Aenean et urna et elit sollicitudin sollicitudin. Nunc vitae lorem eget est fermentum luctus aliquet eget massa. Donec eu ornare risus, ut dignissim mauris. Fusce faucibus magna eget suscipit faucibus. Suspendisse sodales scelerisque sapien ut mollis. Fusce vel sodales lorem. Etiam a arcu malesuada, tempus velit sed, blandit mauris. Fusce et turpis dolor.`,
	`Vestibulum feugiat ante eget faucibus malesuada. Phasellus cursus posuere gravida. Mauris non volutpat ante, id eleifend eros. Morbi in elementum est. Vestibulum tempus sagittis pellentesque. Vivamus vehicula mattis suscipit. Suspendisse ultricies congue ligula sed convallis. Donec non neque quam. Nam convallis ultrices tortor, vel ullamcorper massa malesuada eu. Donec lobortis tortor et dui vehicula iaculis.`,
	`Sed urna tortor, consequat a consequat sit amet, sodales in turpis. Mauris massa purus, finibus sed aliquet id, consequat vitae urna. Aenean imperdiet elit vitae justo porta, et sagittis ante commodo. Aenean et tempus urna. Sed metus nisl, fermentum at augue eu, ultricies interdum massa. Ut sodales nunc convallis mauris pellentesque suscipit. Nulla purus dolor, porttitor ut efficitur sit amet, mollis sed ligula. Morbi pellentesque est nisi, vel molestie est suscipit ut. Vestibulum aliquet sollicitudin risus. Morbi ut congue sem.`,
	//`Aliquam tincidunt sed sapien id hendrerit. Donec sed tincidunt lorem. Phasellus blandit, eros sit amet aliquam convallis, tellus orci elementum elit, non elementum tellus mauris vitae est. Mauris vel suscipit quam. Ut eleifend neque sed tortor ullamcorper, eget maximus leo rhoncus. Maecenas bibendum aliquet justo nec suscipit. Mauris a tristique enim. Pellentesque feugiat augue at tempus rhoncus. In hac habitasse platea dictumst. Pellentesque vitae blandit dui. Suspendisse elementum est eget fringilla consequat. Vestibulum aliquam rhoncus faucibus. Quisque gravida, eros eu aliquam suscipit, turpis quam sollicitudin urna, eget convallis mi ex et ante. Sed posuere vitae velit nec suscipit.`,
}

func main() {
	var result map[string]int
	result = make(map[string]int)
	var wg sync.WaitGroup

	lineCh := reader(data)

	wordsCh1 := lineSplitter(lineCh)
	wordsCh2 := lineSplitter(lineCh)

	wg.Add(1)
	go wordCounter(&result, wordsCh1, wordsCh2, &wg)
	wg.Wait()

	for k, v := range result {
		fmt.Println(v, "->", k)
	}
}

func reader(d []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, line := range d {
			out <- line
		}
		close(out)
	}()
	return out
}

func lineSplitter(lineCh <-chan string) <-chan []string {
	out := make(chan []string)
	go func() {
		for l := range lineCh {
			out <- strings.Split(l, " ")
		}
		close(out)
	}()

	return out
}

func wordCounter(result *map[string]int, wordsCh1 <-chan []string, wordsCh2 <-chan []string, wg *sync.WaitGroup) {
	inCh := make(chan []string)

	var myWg sync.WaitGroup

	reader := func(in <-chan []string, myWg *sync.WaitGroup) {
		for v := range in {		
			inCh <- v
		}
		myWg.Done()
	}

	myWg.Add(1)
	go reader(wordsCh1, &myWg)
	myWg.Add(1)
	go reader(wordsCh2, &myWg)

	go func() {
		for words := range inCh {
			for _, w := range words {
				(*result)[w] = (*result)[w] + 1
			}
		}
	}()

	myWg.Wait()

	close(inCh)

	wg.Done()
}
