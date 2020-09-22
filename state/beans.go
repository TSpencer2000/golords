package state

import (
  "fmt"
)

var createBeanStatement string = `INSERT INTO beans(serverID, userID, amount) VALUES ('%s', '%s', %d)`
var getBeanRowStatement string = `SELECT amount FROM beans WHERE serverID='%s' AND userID='%s'`
var updateBeanRowStatement string = `UPDATE beans SET amount=%d WHERE serverID='%s' AND userID='%s'`
var getTopBeanRowStatement string = `SELECT userID, amount FROM beans WHERE serverID='%s' ORDER BY amount DESC LIMIT %d`

func GetBeansForUser(server, user string) (int, error) {
    var amount int;
    rows, err := database.Query(fmt.Sprintf(getBeanRowStatement, server, user))
    if err != nil {
      return 0, err
    }
    defer rows.Close()
    for rows.Next() {
      err := rows.Scan(&amount)
      if err != nil {
        return 0, err
      }
    }
    // If we didn't get a result, amount will be 0 so we're gravy
    return amount, nil
}

func GetTopNBeans(server string, n int) (map[string]int, error) {
  var user string
  var amount int
  result := make(map[string]int, 0)
  rows, err := database.Query(fmt.Sprintf(getTopBeanRowStatement, server, n))
  if err != nil {
    return result, err
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&user, &amount)
    if err != nil {
      return result, err
    }
    result[user] = amount
  }
  return result, nil
}

func UpdateBeans(server, user string, amount int) (int, error) {
  fmt.Println("Adding %v beans to %v", amount, user)
  var currentScore int;
  var updatedScore int;
  rows, err := database.Query(fmt.Sprintf(getBeanRowStatement, server, user))
  if err != nil {
    return 0, err
  }
  defer rows.Close()
  didGetResult := false
  for rows.Next() {
    err := rows.Scan(&currentScore)
    if err != nil {
      return 0, err
    }
    didGetResult = true
  }
  // Create user if there wasn't a row
  if !didGetResult {
    updatedScore = amount
    err := bbCreateUser(server, user, updatedScore)
    if err != nil {
      return 0, err
    }
    // Otherwise, update the row
  } else {
    updatedScore = currentScore + amount
    // Update the row
    _, err := database.Exec(fmt.Sprintf(updateBeanRowStatement, updatedScore, server, user))
    if err != nil {
      return 0, err
    }
  }
  return updatedScore, nil
}

func bbCreateUser(server, user string, amount int) error {
  _, err := database.Exec(fmt.Sprintf(createBeanStatement, server, user, amount))
  return err
}