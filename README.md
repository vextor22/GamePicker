# GamePicker

GamePicker is a project to create a website which allows users to randomly select a game from their library to play. Planned features include:

- Option to only suggest unplayed or lightly played games.
- Option to add multiple users, and only suggest common games.
- Ability to include or exclude specific games from library.
- Ability to pull user's library even if user has configured a custom SteamID.

## To Run This Project

This project has a comprehensive Docker Compose. The only configuration required at runtime is a config.yml file for the Golang application.

A template is provided, simply fill in your Steam API key, and update the mongo_uri if running outside of the provided Docker Compose environment.
