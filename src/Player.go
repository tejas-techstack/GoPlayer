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
        playlist.removeSong(newSong)
      }
      temploc = temploc.next
    }
    temploc.next = newNode
    newNode.prev = temploc
  }
}

func (pl *pl_Struct) removeSong(remSong string){
  temploc := pl
  // case 4: playlist is empty, edit: Handled in mainloop
  // fmt.Println("Passed Case 4")

  // case 1.a: first song and only song
  if temploc.next==nil && temploc.songName==remSong{
    playlist = nil
    return
  }

  // case 1.b: first song and there are songs after that
  if temploc.next != nil && temploc.songName==remSong{
    playlist = temploc.next
    temploc.next.prev = nil
    return
  }

  // fmt.Println("Passed Case 1.a and 1.b")

  for temploc.next != nil {
    //fmt.Println("Entered For")
    if (temploc.songName == remSong){
        //fmt.Println("Entered if")
        // case 2: middle song
        (temploc.prev).next = temploc.next
        (temploc.next).prev = temploc.prev
        return
    } else {
      //fmt.Println("Entered else")
      temploc = temploc.next
    }
  }
  
  // fmt.Println("Passed Case 2")

  //case 3: last song
  if (temploc.songName == remSong){
    (temploc.prev).next = nil
  } else {
    //case 5 : Song is not there in playlist
    fmt.Println("Song not found in playlist")
  }

  // fmt.Println("Passed Case 3 and 5")
}

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

func (pl *pl_Struct) clear(){
  playlist = nil
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
      case 1:
        menu()
      case 2:
        fmt.Printf("Enter Song to add:")
        var temp string
        fmt.Scanf("%v", &temp)
        playlist.addSong(temp)
      case 3:
        if (playlist == nil){
          fmt.Println("Playlist is empty")
        } else {
          fmt.Printf("Enter Song to remove:")
          var temp string
          fmt.Scanf("%v", &temp)
          playlist.removeSong(temp)
        }
      case 4:
        playlist.show()
      case 5:
        playlist.play()
      case 6:
        playlist.clear()
      case 7:
        fmt.Println("Exiting :)")
        loopVar = false
      case 99:
        playlist.addSong("Test1")
        playlist.addSong("Test2")
        playlist.addSong("Test3")
      default: 
        fmt.Println("Not implemented")
    }
  }
}

func menu(){
  fmt.Printf("######################\n" +
    "\t1. Menu\n" +
    "\t2. Add Song\n" +
    "\t3. Remove Song\n" +
    "\t4. Show Playlist\n" +
    "\t5. Play Songs\n"  +
    "\t6. Clear Playlist\n" +
    "\t7. Exit Player\n"  )

}

func main(){
  mainloop()
}
