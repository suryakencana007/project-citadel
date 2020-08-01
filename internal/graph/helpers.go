package graph

import (
	"context"
	"database/sql"

	"github.com/jordanknott/project-citadel/internal/db"
)

func GetOwnedList(ctx context.Context, r db.Repository, user db.UserAccount) (*OwnedList, error) {
	return &OwnedList{}, nil
}
func GetMemberList(ctx context.Context, r db.Repository, user db.UserAccount) (*MemberList, error) {
	projectMemberIDs, err := r.GetMemberProjectIDsForUserID(ctx, user.UserID)
	if err != sql.ErrNoRows && err != nil {
		return &MemberList{}, err
	}
	var projects []db.Project
	for _, projectID := range projectMemberIDs {
		project, err := r.GetProjectByID(ctx, projectID)
		if err != nil {
			return &MemberList{}, err
		}
		projects = append(projects, project)
	}
	teamMemberIDs, err := r.GetMemberTeamIDsForUserID(ctx, user.UserID)
	if err != sql.ErrNoRows && err != nil {
		return &MemberList{}, err
	}
	var teams []db.Team
	for _, teamID := range teamMemberIDs {
		team, err := r.GetTeamByID(ctx, teamID)
		if err != nil {
			return &MemberList{}, err
		}
		teams = append(teams, team)
	}

	return &MemberList{Teams: teams, Projects: projects}, nil
}
