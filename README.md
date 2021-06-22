# Mattermost Random Meeting

Mattermost plugin to create random meetings. The plugin tries to avoid users you have already talked to.

## Installation

Download the latest version from [releases](https://github.com/juanfran/mattermost-random-meeting/releases).

Go to **System Console > Plugins > Management** upload and enable the plugin.

## Settings 

- **Recurrence** - daily, weekly or monthly meetings.
- **Initial text** - The text that will be send to the users when is time to chat.
- **Number of users per meeting** - Minimum users per meeting.
- **MeetingRooms** - urls separated by commas added to every meeting start message.

## Usage

In any channel you can use the following command. By default you are not going to participate in any meeting until you type `/random-meeting on`.

- `/random-meeting on` - You are available to meet, you have to wait until the the plugin assign you an user group.
- `/random-meeting off` - You don't want to participate in the next recurring meetings.
- `/random-meeting pause` - Toggle meeting participation.

## Admin commands

- `/random-meeting info` - List users that are using the `random-meeting` plugin.
- `/random-meeting add @mention` - Add user.
- `/random-meeting remove @mention` - Remove user.
- `/random-meeting meetings` - Print a JSON string with the previous meetings
- `/random-meeting set_meetings [["user1", "bruce", "martha"]]` - Set the meetings that have are already happened.
