package main

import (
	"errors"

	"github.com/panjf2000/ants/v2"

	sdk "gitee.com/openeuler/go-gitee/gitee"
	libconfig "github.com/opensourceways/community-robot-lib/config"
	libplugin "github.com/opensourceways/community-robot-lib/giteeplugin"
	"github.com/sirupsen/logrus"
)

// TODO: set botName
const botName = ""

type iClient interface {
	GetRepos(org string) ([]sdk.Project, error)
	GetPathContent(org, repo, path, ref string) (sdk.Content, error)
	GetDirectoryTree(org, repo, sha string, recursive int32) (sdk.Tree, error)
	GetRepoAllBranch(org, repo string) ([]sdk.Branch, error)
	GetGiteeRepo(org, repo string) (sdk.Project, error)
	CreateBranch(org, repo, branch, parentBranch string) error
	SetProtectionBranch(org, repo, branch string) error
	CancelProtectionBranch(org, repo, branch string) error
	RemoveRepoMember(org, repo, login string) error
	AddRepoMember(org, repo, login, permission string) error
	CreateRepo(org string, repo sdk.RepositoryPostParam) error
	SetRepoReviewer(org, repo string, reviewer sdk.SetRepoReviewer) error
	UpdateRepo(org, repo string, info sdk.RepoPatchParam) error
}

func newRobot(cli iClient, pool *ants.Pool) *robot {
	return &robot{cli, pool}
}

type robot struct {
	cli  iClient
	pool *ants.Pool
}

func (bot *robot) NewPluginConfig() libconfig.PluginConfig {
	return &configuration{}
}

func (bot *robot) getConfig(cfg libconfig.PluginConfig) (*configuration, error) {
	if c, ok := cfg.(*configuration); ok {
		return c, nil
	}
	return nil, errors.New("can't convert to configuration")
}

func (bot *robot) RegisterEventHandler(p libplugin.HandlerRegitster) {
	p.RegisterIssueHandler(bot.handleIssueEvent)
	p.RegisterPullRequestHandler(bot.handlePREvent)
	p.RegisterNoteEventHandler(bot.handleNoteEvent)
	p.RegisterPushEventHandler(bot.handlePushEvent)
}

func (bot *robot) handlePREvent(e *sdk.PullRequestEvent, cfg libconfig.PluginConfig, log *logrus.Entry) error {
	// TODO: if it doesn't needd to hand PR event, delete this function.
	return nil
}

func (bot *robot) handleIssueEvent(e *sdk.IssueEvent, cfg libconfig.PluginConfig, log *logrus.Entry) error {
	// TODO: if it doesn't needd to hand Issue event, delete this function.
	return nil
}

func (bot *robot) handlePushEvent(e *sdk.PushEvent, cfg libconfig.PluginConfig, log *logrus.Entry) error {
	// TODO: if it doesn't needd to hand Push event, delete this function.
	return nil
}

func (bot *robot) handleNoteEvent(e *sdk.NoteEvent, cfg libconfig.PluginConfig, log *logrus.Entry) error {
	// TODO: if it doesn't needd to hand Note event, delete this function.
	return nil
}
