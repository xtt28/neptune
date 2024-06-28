# Welcome to Neptune's source repository.

Built from the ground up in Go for Minecraft Bedrock Edition, Neptune is a
unique twist on the classic competitive FFA minigame. This repository contains
all source code for the Neptune server, but does not include the server's world.

## Get started with your own instance

1. Clone the repository by clicking the green Code button above the file list.
2. Edit config.toml to your liking.
3. Add your build to the instancedata folder with the name "world". This must be
   a folder containing a Bedrock format world.
4. Edit neptune.toml with the coordinates of the necessary points in your
   server's world. DmgThreshold is the highest Y-value that a player can be at to
   receive damage.

## Features of Neptune

### Gameplay

- To begin gameplay, players choose from a set of kits that come with advantages
  depending on the player's preferred techniques of combat.
- As a reward for kills, players receive bits, the in-game currency. These can
  be exchanged for items and perks.
- All statistics are stored in the database. Players can view them in the
  sidebar HUD and with the `/stats` command.
- Vanilla features such as food depletion and item dropping are disabled.

### Moderation

- Moderators and above can kick players.
- They can issue temporary and permanent bans and mutes with the `/ban` and `/mute` commands.
- Each incident is logged, including the time, type of punishment, duration of
  punishment, reason, subject player, and issuing moderator.
- With the /history command, moderators can easily view the punishment history
  of a given player.

### Permissions

- Player permissions are represented with levels: each level is associated with
  a rank such as administrator or moderator.
- Permissions can be modified by admins with the `/rank set` command.

### Chat

- Chat is formatted based on user rank.
- A powerful profanity filter complements the built-in one provided by Minecraft.
- Players can exchange private messages via the `/msg` command.

## Technological details

- The server is implemented using the Dragonfly library, available at
  <https://github.com/df-mc/dragonfly>.
- GORM (<https://gorm.io>) is used for all database operations with SQLite 3.

## License

All contents of this repository, unless explicitly stated otherwise, are
licensed under the GNU Affero General Public License, v3.0 or later.
