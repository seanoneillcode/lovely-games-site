# lovely-games-site
A site for hosting and playing lovely games.

### Making a change and deploying it
* Ensure correct user details:
  * email
  * ssh key file
  * identity is loaded correctly in current terminal session with:
    ```
    ssh-add -D
    ```
* Create branch
* Commit changes
* Push to Github repository
* Merge branch into main


### Things to be done
* New page to view a single game
* Link each game on the games list to view the single game
* Upload screenshots
* Use a database to store games and data
* Add page to edit a game
  * change details with form
  * delete game button
  * update game state button - release
* filter the main list of games to only be released games
* add new page to manage games - list all games
* add login for admin
* restrict some pages to admin
* restrict some handlers to admin
