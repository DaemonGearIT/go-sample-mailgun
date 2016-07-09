package main

import (
  "github.com/mailgun/mailgun-go"
  "fmt"
)


func main() {
  gun := mailgun.NewMailgun(
    "sandboxf362bedce76d4f988fd0c5dd96e27929.mailgun.org",
    "key-2d333b98c933ed188b192491bb3fb54c",
    "",
  )

  m := mailgun.NewMessage(
    "Sender <sender@example.com>",
    "Subject",
    "Message Body",
    "Gonzalo Bahamondez <gonzalo.bahamondez.c@gmail.com>",
  )

  response, id, _ := gun.Send(m)
  fmt.Printf("Response ID: %s\n", id)
  fmt.Printf("Message from server: %s\n", response)
}


