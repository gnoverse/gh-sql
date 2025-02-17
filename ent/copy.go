// Code generated by ent, DO NOT EDIT.

package ent

// CopyIssue allows to create a new Issue copying the existing
// values of input.
func (ic *IssueCreate) CopyIssue(input *Issue) *IssueCreate {
	ic.SetID(input.ID)
	ic.SetNodeID(input.NodeID)
	ic.SetURL(input.URL)
	ic.SetRepositoryURL(input.RepositoryURL)
	ic.SetLabelsURL(input.LabelsURL)
	ic.SetCommentsURL(input.CommentsURL)
	ic.SetEventsURL(input.EventsURL)
	ic.SetHTMLURL(input.HTMLURL)
	ic.SetNumber(input.Number)
	ic.SetState(input.State)
	if input.StateReason != nil {
		ic.SetStateReason(*input.StateReason)
	}
	ic.SetTitle(input.Title)
	if input.Body != nil {
		ic.SetBody(*input.Body)
	}
	ic.SetLocked(input.Locked)
	if input.ActiveLockReason != nil {
		ic.SetActiveLockReason(*input.ActiveLockReason)
	}
	ic.SetCommentsCount(input.CommentsCount)
	if input.ClosedAt != nil {
		ic.SetClosedAt(*input.ClosedAt)
	}
	ic.SetCreatedAt(input.CreatedAt)
	ic.SetUpdatedAt(input.UpdatedAt)
	ic.SetDraft(input.Draft)
	ic.SetAuthorAssociation(input.AuthorAssociation)
	ic.SetReactions(input.Reactions)
	return ic
}

// CopyIssue allows to update a Issue copying the existing
// values of input.
func (iuo *IssueUpdateOne) CopyIssue(input *Issue) *IssueUpdateOne {
	iuo.SetNodeID(input.NodeID)
	iuo.SetURL(input.URL)
	iuo.SetRepositoryURL(input.RepositoryURL)
	iuo.SetLabelsURL(input.LabelsURL)
	iuo.SetCommentsURL(input.CommentsURL)
	iuo.SetEventsURL(input.EventsURL)
	iuo.SetHTMLURL(input.HTMLURL)
	iuo.SetNumber(input.Number)
	iuo.SetState(input.State)
	if input.StateReason != nil {
		iuo.SetStateReason(*input.StateReason)
	}
	iuo.SetTitle(input.Title)
	if input.Body != nil {
		iuo.SetBody(*input.Body)
	}
	iuo.SetLocked(input.Locked)
	if input.ActiveLockReason != nil {
		iuo.SetActiveLockReason(*input.ActiveLockReason)
	}
	iuo.SetCommentsCount(input.CommentsCount)
	if input.ClosedAt != nil {
		iuo.SetClosedAt(*input.ClosedAt)
	}
	iuo.SetCreatedAt(input.CreatedAt)
	iuo.SetUpdatedAt(input.UpdatedAt)
	iuo.SetDraft(input.Draft)
	iuo.SetAuthorAssociation(input.AuthorAssociation)
	iuo.SetReactions(input.Reactions)
	return iuo
}

// CopyIssueComment allows to create a new IssueComment copying the existing
// values of input.
func (icc *IssueCommentCreate) CopyIssueComment(input *IssueComment) *IssueCommentCreate {
	icc.SetID(input.ID)
	icc.SetNodeID(input.NodeID)
	icc.SetURL(input.URL)
	icc.SetBody(input.Body)
	icc.SetHTMLURL(input.HTMLURL)
	icc.SetCreatedAt(input.CreatedAt)
	icc.SetUpdatedAt(input.UpdatedAt)
	icc.SetIssueURL(input.IssueURL)
	icc.SetAuthorAssociation(input.AuthorAssociation)
	icc.SetReactions(input.Reactions)
	return icc
}

// CopyIssueComment allows to update a IssueComment copying the existing
// values of input.
func (icuo *IssueCommentUpdateOne) CopyIssueComment(input *IssueComment) *IssueCommentUpdateOne {
	icuo.SetNodeID(input.NodeID)
	icuo.SetURL(input.URL)
	icuo.SetBody(input.Body)
	icuo.SetHTMLURL(input.HTMLURL)
	icuo.SetCreatedAt(input.CreatedAt)
	icuo.SetUpdatedAt(input.UpdatedAt)
	icuo.SetIssueURL(input.IssueURL)
	icuo.SetAuthorAssociation(input.AuthorAssociation)
	icuo.SetReactions(input.Reactions)
	return icuo
}

// CopyPullRequest allows to create a new PullRequest copying the existing
// values of input.
func (prc *PullRequestCreate) CopyPullRequest(input *PullRequest) *PullRequestCreate {
	prc.SetID(input.ID)
	prc.SetURL(input.URL)
	prc.SetNodeID(input.NodeID)
	prc.SetHTMLURL(input.HTMLURL)
	prc.SetDiffURL(input.DiffURL)
	prc.SetPatchURL(input.PatchURL)
	prc.SetIssueURL(input.IssueURL)
	prc.SetCommitsURL(input.CommitsURL)
	prc.SetReviewCommentsURL(input.ReviewCommentsURL)
	prc.SetReviewCommentURL(input.ReviewCommentURL)
	prc.SetCommentsURL(input.CommentsURL)
	prc.SetStatusesURL(input.StatusesURL)
	prc.SetNumber(input.Number)
	prc.SetState(input.State)
	prc.SetLocked(input.Locked)
	prc.SetTitle(input.Title)
	prc.SetBody(input.Body)
	prc.SetCreatedAt(input.CreatedAt)
	prc.SetUpdatedAt(input.UpdatedAt)
	prc.SetClosedAt(input.ClosedAt)
	prc.SetMergedAt(input.MergedAt)
	prc.SetMergeCommitSha(input.MergeCommitSha)
	prc.SetHead(input.Head)
	prc.SetBase(input.Base)
	if input.Draft != nil {
		prc.SetDraft(*input.Draft)
	}
	prc.SetAuthorAssociation(input.AuthorAssociation)
	return prc
}

// CopyPullRequest allows to update a PullRequest copying the existing
// values of input.
func (pruo *PullRequestUpdateOne) CopyPullRequest(input *PullRequest) *PullRequestUpdateOne {
	pruo.SetURL(input.URL)
	pruo.SetNodeID(input.NodeID)
	pruo.SetHTMLURL(input.HTMLURL)
	pruo.SetDiffURL(input.DiffURL)
	pruo.SetPatchURL(input.PatchURL)
	pruo.SetIssueURL(input.IssueURL)
	pruo.SetCommitsURL(input.CommitsURL)
	pruo.SetReviewCommentsURL(input.ReviewCommentsURL)
	pruo.SetReviewCommentURL(input.ReviewCommentURL)
	pruo.SetCommentsURL(input.CommentsURL)
	pruo.SetStatusesURL(input.StatusesURL)
	pruo.SetNumber(input.Number)
	pruo.SetState(input.State)
	pruo.SetLocked(input.Locked)
	pruo.SetTitle(input.Title)
	pruo.SetBody(input.Body)
	pruo.SetCreatedAt(input.CreatedAt)
	pruo.SetUpdatedAt(input.UpdatedAt)
	pruo.SetClosedAt(input.ClosedAt)
	pruo.SetMergedAt(input.MergedAt)
	pruo.SetMergeCommitSha(input.MergeCommitSha)
	pruo.SetHead(input.Head)
	pruo.SetBase(input.Base)
	if input.Draft != nil {
		pruo.SetDraft(*input.Draft)
	}
	pruo.SetAuthorAssociation(input.AuthorAssociation)
	return pruo
}

// CopyRepository allows to create a new Repository copying the existing
// values of input.
func (rc *RepositoryCreate) CopyRepository(input *Repository) *RepositoryCreate {
	rc.SetID(input.ID)
	rc.SetNodeID(input.NodeID)
	rc.SetName(input.Name)
	rc.SetFullName(input.FullName)
	rc.SetPrivate(input.Private)
	rc.SetHTMLURL(input.HTMLURL)
	rc.SetDescription(input.Description)
	rc.SetFork(input.Fork)
	rc.SetURL(input.URL)
	rc.SetArchiveURL(input.ArchiveURL)
	rc.SetAssigneesURL(input.AssigneesURL)
	rc.SetBlobsURL(input.BlobsURL)
	rc.SetBranchesURL(input.BranchesURL)
	rc.SetCollaboratorsURL(input.CollaboratorsURL)
	rc.SetCommentsURL(input.CommentsURL)
	rc.SetCommitsURL(input.CommitsURL)
	rc.SetCompareURL(input.CompareURL)
	rc.SetContentsURL(input.ContentsURL)
	rc.SetContributorsURL(input.ContributorsURL)
	rc.SetDeploymentsURL(input.DeploymentsURL)
	rc.SetDownloadsURL(input.DownloadsURL)
	rc.SetEventsURL(input.EventsURL)
	rc.SetForksURL(input.ForksURL)
	rc.SetGitCommitsURL(input.GitCommitsURL)
	rc.SetGitRefsURL(input.GitRefsURL)
	rc.SetGitTagsURL(input.GitTagsURL)
	rc.SetGitURL(input.GitURL)
	rc.SetIssueCommentURL(input.IssueCommentURL)
	rc.SetIssueEventsURL(input.IssueEventsURL)
	rc.SetIssuesURL(input.IssuesURL)
	rc.SetKeysURL(input.KeysURL)
	rc.SetLabelsURL(input.LabelsURL)
	rc.SetLanguagesURL(input.LanguagesURL)
	rc.SetMergesURL(input.MergesURL)
	rc.SetMilestonesURL(input.MilestonesURL)
	rc.SetNotificationsURL(input.NotificationsURL)
	rc.SetPullsURL(input.PullsURL)
	rc.SetReleasesURL(input.ReleasesURL)
	rc.SetSSHURL(input.SSHURL)
	rc.SetStargazersURL(input.StargazersURL)
	rc.SetStatusesURL(input.StatusesURL)
	rc.SetSubscribersURL(input.SubscribersURL)
	rc.SetSubscriptionURL(input.SubscriptionURL)
	rc.SetTagsURL(input.TagsURL)
	rc.SetTeamsURL(input.TeamsURL)
	rc.SetTreesURL(input.TreesURL)
	rc.SetCloneURL(input.CloneURL)
	if input.MirrorURL != nil {
		rc.SetMirrorURL(*input.MirrorURL)
	}
	rc.SetHooksURL(input.HooksURL)
	rc.SetSvnURL(input.SvnURL)
	rc.SetHomepage(input.Homepage)
	rc.SetLanguage(input.Language)
	rc.SetForksCount(input.ForksCount)
	rc.SetStargazersCount(input.StargazersCount)
	rc.SetWatchersCount(input.WatchersCount)
	rc.SetSize(input.Size)
	rc.SetDefaultBranch(input.DefaultBranch)
	rc.SetOpenIssuesCount(input.OpenIssuesCount)
	rc.SetIsTemplate(input.IsTemplate)
	rc.SetTopics(input.Topics)
	rc.SetHasIssuesEnabled(input.HasIssuesEnabled)
	rc.SetHasProjects(input.HasProjects)
	rc.SetHasWiki(input.HasWiki)
	rc.SetHasPages(input.HasPages)
	rc.SetHasDownloads(input.HasDownloads)
	rc.SetHasDiscussions(input.HasDiscussions)
	rc.SetArchived(input.Archived)
	rc.SetDisabled(input.Disabled)
	if input.Visibility != nil {
		rc.SetVisibility(*input.Visibility)
	}
	rc.SetPushedAt(input.PushedAt)
	rc.SetCreatedAt(input.CreatedAt)
	rc.SetUpdatedAt(input.UpdatedAt)
	rc.SetSubscribersCount(input.SubscribersCount)
	rc.SetNetworkCount(input.NetworkCount)
	rc.SetForks(input.Forks)
	rc.SetOpenIssues(input.OpenIssues)
	rc.SetWatchers(input.Watchers)
	rc.SetLicense(input.License)
	return rc
}

// CopyRepository allows to update a Repository copying the existing
// values of input.
func (ruo *RepositoryUpdateOne) CopyRepository(input *Repository) *RepositoryUpdateOne {
	ruo.SetNodeID(input.NodeID)
	ruo.SetName(input.Name)
	ruo.SetFullName(input.FullName)
	ruo.SetPrivate(input.Private)
	ruo.SetHTMLURL(input.HTMLURL)
	ruo.SetDescription(input.Description)
	ruo.SetFork(input.Fork)
	ruo.SetURL(input.URL)
	ruo.SetArchiveURL(input.ArchiveURL)
	ruo.SetAssigneesURL(input.AssigneesURL)
	ruo.SetBlobsURL(input.BlobsURL)
	ruo.SetBranchesURL(input.BranchesURL)
	ruo.SetCollaboratorsURL(input.CollaboratorsURL)
	ruo.SetCommentsURL(input.CommentsURL)
	ruo.SetCommitsURL(input.CommitsURL)
	ruo.SetCompareURL(input.CompareURL)
	ruo.SetContentsURL(input.ContentsURL)
	ruo.SetContributorsURL(input.ContributorsURL)
	ruo.SetDeploymentsURL(input.DeploymentsURL)
	ruo.SetDownloadsURL(input.DownloadsURL)
	ruo.SetEventsURL(input.EventsURL)
	ruo.SetForksURL(input.ForksURL)
	ruo.SetGitCommitsURL(input.GitCommitsURL)
	ruo.SetGitRefsURL(input.GitRefsURL)
	ruo.SetGitTagsURL(input.GitTagsURL)
	ruo.SetGitURL(input.GitURL)
	ruo.SetIssueCommentURL(input.IssueCommentURL)
	ruo.SetIssueEventsURL(input.IssueEventsURL)
	ruo.SetIssuesURL(input.IssuesURL)
	ruo.SetKeysURL(input.KeysURL)
	ruo.SetLabelsURL(input.LabelsURL)
	ruo.SetLanguagesURL(input.LanguagesURL)
	ruo.SetMergesURL(input.MergesURL)
	ruo.SetMilestonesURL(input.MilestonesURL)
	ruo.SetNotificationsURL(input.NotificationsURL)
	ruo.SetPullsURL(input.PullsURL)
	ruo.SetReleasesURL(input.ReleasesURL)
	ruo.SetSSHURL(input.SSHURL)
	ruo.SetStargazersURL(input.StargazersURL)
	ruo.SetStatusesURL(input.StatusesURL)
	ruo.SetSubscribersURL(input.SubscribersURL)
	ruo.SetSubscriptionURL(input.SubscriptionURL)
	ruo.SetTagsURL(input.TagsURL)
	ruo.SetTeamsURL(input.TeamsURL)
	ruo.SetTreesURL(input.TreesURL)
	ruo.SetCloneURL(input.CloneURL)
	if input.MirrorURL != nil {
		ruo.SetMirrorURL(*input.MirrorURL)
	}
	ruo.SetHooksURL(input.HooksURL)
	ruo.SetSvnURL(input.SvnURL)
	ruo.SetHomepage(input.Homepage)
	ruo.SetLanguage(input.Language)
	ruo.SetForksCount(input.ForksCount)
	ruo.SetStargazersCount(input.StargazersCount)
	ruo.SetWatchersCount(input.WatchersCount)
	ruo.SetSize(input.Size)
	ruo.SetDefaultBranch(input.DefaultBranch)
	ruo.SetOpenIssuesCount(input.OpenIssuesCount)
	ruo.SetIsTemplate(input.IsTemplate)
	ruo.SetTopics(input.Topics)
	ruo.SetHasIssuesEnabled(input.HasIssuesEnabled)
	ruo.SetHasProjects(input.HasProjects)
	ruo.SetHasWiki(input.HasWiki)
	ruo.SetHasPages(input.HasPages)
	ruo.SetHasDownloads(input.HasDownloads)
	ruo.SetHasDiscussions(input.HasDiscussions)
	ruo.SetArchived(input.Archived)
	ruo.SetDisabled(input.Disabled)
	if input.Visibility != nil {
		ruo.SetVisibility(*input.Visibility)
	}
	ruo.SetPushedAt(input.PushedAt)
	ruo.SetCreatedAt(input.CreatedAt)
	ruo.SetUpdatedAt(input.UpdatedAt)
	ruo.SetSubscribersCount(input.SubscribersCount)
	ruo.SetNetworkCount(input.NetworkCount)
	ruo.SetForks(input.Forks)
	ruo.SetOpenIssues(input.OpenIssues)
	ruo.SetWatchers(input.Watchers)
	ruo.SetLicense(input.License)
	return ruo
}

// CopyTimelineEvent allows to create a new TimelineEvent copying the existing
// values of input.
func (tec *TimelineEventCreate) CopyTimelineEvent(input *TimelineEvent) *TimelineEventCreate {
	tec.SetNumericID(input.NumericID)
	tec.SetNodeID(input.NodeID)
	tec.SetURL(input.URL)
	tec.SetEvent(input.Event)
	if input.CommitID != nil {
		tec.SetCommitID(*input.CommitID)
	}
	if input.CommitURL != nil {
		tec.SetCommitURL(*input.CommitURL)
	}
	tec.SetCreatedAt(input.CreatedAt)
	tec.SetData(input.Data)
	return tec
}

// CopyTimelineEvent allows to update a TimelineEvent copying the existing
// values of input.
func (teuo *TimelineEventUpdateOne) CopyTimelineEvent(input *TimelineEvent) *TimelineEventUpdateOne {
	teuo.SetNumericID(input.NumericID)
	teuo.SetNodeID(input.NodeID)
	teuo.SetURL(input.URL)
	teuo.SetEvent(input.Event)
	if input.CommitID != nil {
		teuo.SetCommitID(*input.CommitID)
	}
	if input.CommitURL != nil {
		teuo.SetCommitURL(*input.CommitURL)
	}
	teuo.SetCreatedAt(input.CreatedAt)
	teuo.SetData(input.Data)
	return teuo
}

// CopyUser allows to create a new User copying the existing
// values of input.
func (uc *UserCreate) CopyUser(input *User) *UserCreate {
	uc.SetID(input.ID)
	uc.SetLogin(input.Login)
	uc.SetNodeID(input.NodeID)
	uc.SetAvatarURL(input.AvatarURL)
	uc.SetGravatarID(input.GravatarID)
	uc.SetURL(input.URL)
	uc.SetHTMLURL(input.HTMLURL)
	uc.SetFollowersURL(input.FollowersURL)
	uc.SetFollowingURL(input.FollowingURL)
	uc.SetGistsURL(input.GistsURL)
	uc.SetStarredURL(input.StarredURL)
	uc.SetSubscriptionsURL(input.SubscriptionsURL)
	uc.SetOrganizationsURL(input.OrganizationsURL)
	uc.SetReposURL(input.ReposURL)
	uc.SetEventsURL(input.EventsURL)
	uc.SetReceivedEventsURL(input.ReceivedEventsURL)
	uc.SetType(input.Type)
	uc.SetSiteAdmin(input.SiteAdmin)
	uc.SetName(input.Name)
	uc.SetCompany(input.Company)
	uc.SetBlog(input.Blog)
	uc.SetLocation(input.Location)
	uc.SetEmail(input.Email)
	uc.SetHireable(input.Hireable)
	uc.SetBio(input.Bio)
	uc.SetPublicRepos(input.PublicRepos)
	uc.SetPublicGists(input.PublicGists)
	uc.SetFollowers(input.Followers)
	uc.SetFollowing(input.Following)
	uc.SetCreatedAt(input.CreatedAt)
	uc.SetUpdatedAt(input.UpdatedAt)
	return uc
}

// CopyUser allows to update a User copying the existing
// values of input.
func (uuo *UserUpdateOne) CopyUser(input *User) *UserUpdateOne {
	uuo.SetLogin(input.Login)
	uuo.SetNodeID(input.NodeID)
	uuo.SetAvatarURL(input.AvatarURL)
	uuo.SetGravatarID(input.GravatarID)
	uuo.SetURL(input.URL)
	uuo.SetHTMLURL(input.HTMLURL)
	uuo.SetFollowersURL(input.FollowersURL)
	uuo.SetFollowingURL(input.FollowingURL)
	uuo.SetGistsURL(input.GistsURL)
	uuo.SetStarredURL(input.StarredURL)
	uuo.SetSubscriptionsURL(input.SubscriptionsURL)
	uuo.SetOrganizationsURL(input.OrganizationsURL)
	uuo.SetReposURL(input.ReposURL)
	uuo.SetEventsURL(input.EventsURL)
	uuo.SetReceivedEventsURL(input.ReceivedEventsURL)
	uuo.SetType(input.Type)
	uuo.SetSiteAdmin(input.SiteAdmin)
	uuo.SetName(input.Name)
	uuo.SetCompany(input.Company)
	uuo.SetBlog(input.Blog)
	uuo.SetLocation(input.Location)
	uuo.SetEmail(input.Email)
	uuo.SetHireable(input.Hireable)
	uuo.SetBio(input.Bio)
	uuo.SetPublicRepos(input.PublicRepos)
	uuo.SetPublicGists(input.PublicGists)
	uuo.SetFollowers(input.Followers)
	uuo.SetFollowing(input.Following)
	uuo.SetCreatedAt(input.CreatedAt)
	uuo.SetUpdatedAt(input.UpdatedAt)
	return uuo
}
