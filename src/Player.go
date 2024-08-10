package main

import(
  "fmt"
  "time"
)


type pl_Struct struct {
  prev       *pl_Struct
  songName   string
  next       *pl_Struct
}

var playlist *pl_Struct

func (pl *pl_Struct) addSong(newSong string) {
  temploc := pl
  newNode := &pl_Struct{songName : newSong}

  if pl == nil {
    playlist = newNode
  } else {
    for temploc.next != nil {
      if (temploc.songName == newSong){
        fmt.Println("Song already in playlist Moving to front")
        // playlist.removeSong(newSong)
      }
      temploc = temploc.next
    }
    temploc.next = newNode
    newNode.prev = temploc
  }
}

/*
func (pl *pl_Struct) removeSong(remSong string){
  temploc := pl
  
  for temploc.next != nil {
    if (temploc.next.songName == newSong){
      
    } else {
      temploc = temploc.next
    }
  }
}
*/

func (pl *pl_Struct) show(){
  if pl == nil{
    fmt.Println("Playlist is empty")
    return
  }
  temploc := pl
  i := 1
  for temploc != nil{
    fmt.Printf("%v%10v\n", i , temploc.songName)
    i++
    temploc = temploc.next
  }
}

func (pl *pl_Struct) play(){
  if pl == nil{
    fmt.Println("Playlist is empty")
    return
  }
  temploc := pl
  for temploc != nil{
    fmt.Println("Playing Song:", temploc.songName)
// -----------to be replaced---------------
    time.Sleep(1*time.Second)
// ----------with actual song--------------
    temploc = temploc.next
  }
}

func mainloop(){
  var userChoice int
  var loopVar = true

  fmt.Println("############   GoPlayer   #############")
  menu()

  for loopVar {
    fmt.Printf("Enter Choice:")
    fmt.Scanf("%v", &userChoice)
    switch userChoice {
      case 0:
        menu()
      case 1:
        fmt.Printf("Enter Song:")
        var temp string
        fmt.Scanf("%v", &temp)
        playlist.addSong(temp)
      case 2:
        playlist.show()
      case 3:
        playlist.play()
      case 4:
        fmt.Println("Exiting :)")
        loopVar = false
      case 5:
        playlist.addSong("Test1")
        playlist.addSong("Test2")
        playlist.addSong("Test3")
      default: 
        fmt.Println("Not implemented")
    }
  }
}

func menu(){
  fmt.Printf("#######################\n" +
    "\t0. Menu\n" +
    "\t1. Add Song\n" +
    "\t2. Show playlist\n" +
    "\t3. Play Songs\n" +
    "\t4. Exit Player\n" +
    "\t5. Add Test Songs\n")
}

func main(){
  mainloop()
}
