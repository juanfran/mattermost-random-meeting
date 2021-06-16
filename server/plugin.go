package main

import (
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/robfig/cron/v3"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	cron        *cron.Cron
	cronEntryID cron.EntryID

	users         []string
	paused        []string
	usersMeetings [][]string

	botUserID string
}

func (p *Plugin) refreshCron(configuration *configuration) {
	p.cron.Remove(p.cronEntryID)
	p.addCronFunc()
}

func (p *Plugin) addCronFunc() {
	config := p.getConfiguration()
	configCron := config.Cron

	if configCron == "" {
		configCron = "@weekly"
	}

	if config.Cron == "custom" && len(config.CustomCron) > 0 {
		configCron = config.CustomCron
	}

	// every minute "* * * * *"
	p.cronEntryID, _ = p.cron.AddFunc(configCron, func() {
		p.runMeetings()
	})
}

func (p *Plugin) runMeetings() {
	users := p.getAvailableUsers()

	// todo
	reverseAny(p.usersMeetings)

	meetings := getMeetingsShuffleUsers(users, p.configuration.NumUsersPerMeeting, p.userMeetingsReversed)

	for _, meeting := range meetings {
		p.startMeeting(meeting)
	}
}

func (p *Plugin) startMeeting(meeting []string) {
	maxMeetings := 50

	p.usersMeetings = append(p.usersMeetings, meeting)

	if len(p.usersMeetings) > maxMeetings {
		_, p.usersMeetings = p.usersMeetings[0], p.usersMeetings[1:]
	}

	users := append(meeting, p.botUserID)

	channel, _ := p.API.GetGroupChannel(users)

	config := p.getConfiguration()

	post := &model.Post{
		UserId:    p.botUserID,
		ChannelId: channel.Id,
		Message:   config.InitText,
	}

	p.persistMeetings()
	p.API.CreatePost(post)
}

func (p *Plugin) UserHasLeftTeam(c *plugin.Context, teamMember *model.TeamMember) {
	p.users = Remove(p.users, teamMember.UserId)
	p.persistMeetings()
}

func (p *Plugin) getAvailableUsers() []string {
	var users []string

	for _, userId := range p.users {
		if !Contains(p.paused, userId) {
			users = append(users, userId)
		}
	}

	return users
}

func (p *Plugin) addUser(userID string) {
	if !Contains(p.users, userID) {
		p.users = append(p.users, userID)
	}
}

func (p *Plugin) removeUser(userID string) {
	p.users = Remove(p.users, userID)
	p.persistMeetings()
}

func (p *Plugin) OnDeactivate() error {
	p.cron.Remove(p.cronEntryID)
	p.cron.Stop()

	return p.persistUsers()
}

func (p *Plugin) usersMeetingsByUsername() [][]string {
	mettings := [][]string{}

	for _, meeting := range p.usersMeetings {
		newMeeting := []string{}

		for _, userId := range meeting {
			userData, _ := p.API.GetUser(userId)

			newMeeting = append(newMeeting, userData.Username)
		}

		mettings = append(mettings, newMeeting)
	}

	return mettings
}
