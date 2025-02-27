package mockstore

import (
	"context"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/search"
	"github.com/grafana/grafana/pkg/services/sqlstore"
)

type SQLStoreMock struct {
	ExpectedUser       *models.User
	ExpectedDatasource *models.DataSource
	ExpectedError      error
}

func NewSQLStoreMock() *SQLStoreMock {
	return &SQLStoreMock{}
}

func (m SQLStoreMock) DeleteExpiredSnapshots(ctx context.Context, cmd *models.DeleteExpiredSnapshotsCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CreateDashboardSnapshot(ctx context.Context, cmd *models.CreateDashboardSnapshotCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteDashboardSnapshot(ctx context.Context, cmd *models.DeleteDashboardSnapshotCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDashboardSnapshot(query *models.GetDashboardSnapshotQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) HasEditPermissionInFolders(ctx context.Context, query *models.HasEditPermissionInFoldersQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SearchDashboardSnapshots(query *models.GetDashboardSnapshotsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetOrgByName(name string) (*models.Org, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) CreateOrgWithMember(name string, userID int64) (models.Org, error) {
	return models.Org{}, nil
}

func (m SQLStoreMock) UpdateOrg(ctx context.Context, cmd *models.UpdateOrgCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateOrgAddress(ctx context.Context, cmd *models.UpdateOrgAddressCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteOrg(ctx context.Context, cmd *models.DeleteOrgCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetProvisionedDataByDashboardID(dashboardID int64) (*models.DashboardProvisioning, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) GetProvisionedDataByDashboardUID(orgID int64, dashboardUID string) (*models.DashboardProvisioning, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) SaveProvisionedDashboard(cmd models.SaveDashboardCommand, provisioning *models.DashboardProvisioning) (*models.Dashboard, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) GetProvisionedDashboardData(name string) ([]*models.DashboardProvisioning, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) DeleteOrphanedProvisionedDashboards(ctx context.Context, cmd *models.DeleteOrphanedProvisionedDashboardsCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CreateLoginAttempt(ctx context.Context, cmd *models.CreateLoginAttemptCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteOldLoginAttempts(ctx context.Context, cmd *models.DeleteOldLoginAttemptsCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CloneUserToServiceAccount(ctx context.Context, siUser *models.SignedInUser) (*models.User, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) CreateServiceAccountForApikey(ctx context.Context, orgId int64, keyname string, role models.RoleType) (*models.User, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) CreateUser(ctx context.Context, cmd models.CreateUserCommand) (*models.User, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) GetUserById(ctx context.Context, query *models.GetUserByIdQuery) error {
	query.Result = m.ExpectedUser
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserByLogin(ctx context.Context, query *models.GetUserByLoginQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserByEmail(ctx context.Context, query *models.GetUserByEmailQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateUser(ctx context.Context, cmd *models.UpdateUserCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) ChangeUserPassword(ctx context.Context, cmd *models.ChangeUserPasswordCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateUserLastSeenAt(ctx context.Context, cmd *models.UpdateUserLastSeenAtCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SetUsingOrg(ctx context.Context, cmd *models.SetUsingOrgCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserProfile(ctx context.Context, query *models.GetUserProfileQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserOrgList(ctx context.Context, query *models.GetUserOrgListQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetSignedInUserWithCacheCtx(ctx context.Context, query *models.GetSignedInUserQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetSignedInUser(ctx context.Context, query *models.GetSignedInUserQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) BatchDisableUsers(ctx context.Context, cmd *models.BatchDisableUsersCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteUser(ctx context.Context, cmd *models.DeleteUserCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateUserPermissions(userID int64, isAdmin bool) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SetUserHelpFlag(ctx context.Context, cmd *models.SetUserHelpFlagCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CreateTeam(name string, email string, orgID int64) (models.Team, error) {
	return models.Team{
		Name:  name,
		Email: email,
		OrgId: orgID,
	}, nil
}

func (m SQLStoreMock) UpdateTeam(ctx context.Context, cmd *models.UpdateTeamCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteTeam(ctx context.Context, cmd *models.DeleteTeamCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SearchTeams(ctx context.Context, query *models.SearchTeamsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetTeamById(ctx context.Context, query *models.GetTeamByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetTeamsByUser(ctx context.Context, query *models.GetTeamsByUserQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) AddTeamMember(userID int64, orgID int64, teamID int64, isExternal bool, permission models.PermissionType) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateTeamMember(ctx context.Context, cmd *models.UpdateTeamMemberCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) IsTeamMember(orgId int64, teamId int64, userId int64) (bool, error) {
	return false, nil
}

func (m SQLStoreMock) RemoveTeamMember(ctx context.Context, cmd *models.RemoveTeamMemberCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetTeamMembers(ctx context.Context, query *models.GetTeamMembersQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) NewSession(ctx context.Context) *sqlstore.DBSession {
	return nil
}

func (m SQLStoreMock) WithDbSession(ctx context.Context, callback sqlstore.DBTransactionFunc) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetPreferencesWithDefaults(ctx context.Context, query *models.GetPreferencesWithDefaultsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetPreferences(ctx context.Context, query *models.GetPreferencesQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SavePreferences(ctx context.Context, cmd *models.SavePreferencesCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetPluginSettings(ctx context.Context, orgID int64) ([]*models.PluginSettingInfoDTO, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) GetPluginSettingById(ctx context.Context, query *models.GetPluginSettingByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdatePluginSetting(ctx context.Context, cmd *models.UpdatePluginSettingCmd) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdatePluginSettingVersion(ctx context.Context, cmd *models.UpdatePluginSettingVersionCmd) error {
	return m.ExpectedError
}

func (m SQLStoreMock) IsStarredByUserCtx(ctx context.Context, query *models.IsStarredByUserQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) StarDashboard(ctx context.Context, cmd *models.StarDashboardCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UnstarDashboard(ctx context.Context, cmd *models.UnstarDashboardCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserStars(ctx context.Context, query *models.GetUserStarsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetOrgQuotaByTarget(ctx context.Context, query *models.GetOrgQuotaByTargetQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetOrgQuotas(ctx context.Context, query *models.GetOrgQuotasQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateOrgQuota(ctx context.Context, cmd *models.UpdateOrgQuotaCmd) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserQuotaByTarget(ctx context.Context, query *models.GetUserQuotaByTargetQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetUserQuotas(ctx context.Context, query *models.GetUserQuotasQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateUserQuota(ctx context.Context, cmd *models.UpdateUserQuotaCmd) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetGlobalQuotaByTarget(ctx context.Context, query *models.GetGlobalQuotaByTargetQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) WithTransactionalDbSession(ctx context.Context, callback sqlstore.DBTransactionFunc) error {
	return m.ExpectedError
}

func (m SQLStoreMock) InTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDashboardVersion(ctx context.Context, query *models.GetDashboardVersionQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDashboardVersions(ctx context.Context, query *models.GetDashboardVersionsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteExpiredVersions(ctx context.Context, cmd *models.DeleteExpiredVersionsCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateDashboardACL(ctx context.Context, dashboardID int64, items []*models.DashboardAcl) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateDashboardACLCtx(ctx context.Context, dashboardID int64, items []*models.DashboardAcl) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDashboardAclInfoList(ctx context.Context, query *models.GetDashboardAclInfoListQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CreatePlaylist(ctx context.Context, cmd *models.CreatePlaylistCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdatePlaylist(ctx context.Context, cmd *models.UpdatePlaylistCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetPlaylist(ctx context.Context, query *models.GetPlaylistByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeletePlaylist(ctx context.Context, cmd *models.DeletePlaylistCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SearchPlaylists(ctx context.Context, query *models.GetPlaylistsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetPlaylistItem(ctx context.Context, query *models.GetPlaylistItemsByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAlertById(ctx context.Context, query *models.GetAlertByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAllAlertQueryHandler(ctx context.Context, query *models.GetAllAlertsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) HandleAlertsQuery(ctx context.Context, query *models.GetAlertsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SaveAlerts(ctx context.Context, dashID int64, alerts []*models.Alert) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SetAlertState(ctx context.Context, cmd *models.SetAlertStateCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) PauseAlert(ctx context.Context, cmd *models.PauseAlertCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) PauseAllAlerts(ctx context.Context, cmd *models.PauseAllAlertCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAlertStatesForDashboard(ctx context.Context, query *models.GetAlertStatesForDashboardQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) AddOrgUser(ctx context.Context, cmd *models.AddOrgUserCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateOrgUser(ctx context.Context, cmd *models.UpdateOrgUserCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetOrgUsers(ctx context.Context, query *models.GetOrgUsersQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SearchOrgUsers(ctx context.Context, query *models.SearchOrgUsersQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) RemoveOrgUser(ctx context.Context, cmd *models.RemoveOrgUserCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SaveDashboard(cmd models.SaveDashboardCommand) (*models.Dashboard, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) GetDashboard(id int64, orgID int64, uid string, slug string) (*models.Dashboard, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) GetFolderByTitle(orgID int64, title string) (*models.Dashboard, error) {
	return nil, m.ExpectedError
}

func (m SQLStoreMock) SearchDashboards(ctx context.Context, query *search.FindPersistedDashboardsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteDashboard(ctx context.Context, cmd *models.DeleteDashboardCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDashboards(ctx context.Context, query *models.GetDashboardsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDashboardUIDById(ctx context.Context, query *models.GetDashboardRefByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) ValidateDashboardBeforeSave(dashboard *models.Dashboard, overwrite bool) (bool, error) {
	return false, nil
}

func (m SQLStoreMock) GetDataSource(ctx context.Context, query *models.GetDataSourceQuery) error {
	query.Result = m.ExpectedDatasource
	return m.ExpectedError
}

func (m SQLStoreMock) GetDataSources(ctx context.Context, query *models.GetDataSourcesQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDataSourcesByType(ctx context.Context, query *models.GetDataSourcesByTypeQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetDefaultDataSource(ctx context.Context, query *models.GetDefaultDataSourceQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteDataSource(ctx context.Context, cmd *models.DeleteDataSourceCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) AddDataSource(ctx context.Context, cmd *models.AddDataSourceCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateDataSource(ctx context.Context, cmd *models.UpdateDataSourceCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) Migrate() error {
	return m.ExpectedError
}

func (m SQLStoreMock) Sync() error {
	return m.ExpectedError
}

func (m SQLStoreMock) Reset() error {
	return m.ExpectedError
}

func (m SQLStoreMock) Quote(value string) string {
	return ""
}

func (m SQLStoreMock) DeleteAlertNotification(ctx context.Context, cmd *models.DeleteAlertNotificationCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) DeleteAlertNotificationWithUid(ctx context.Context, cmd *models.DeleteAlertNotificationWithUidCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAlertNotifications(ctx context.Context, query *models.GetAlertNotificationsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAlertNotificationUidWithId(ctx context.Context, query *models.GetAlertNotificationUidQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAlertNotificationsWithUid(ctx context.Context, query *models.GetAlertNotificationsWithUidQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAllAlertNotifications(ctx context.Context, query *models.GetAllAlertNotificationsQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAlertNotificationsWithUidToSend(ctx context.Context, query *models.GetAlertNotificationsWithUidToSendQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CreateAlertNotificationCommand(ctx context.Context, cmd *models.CreateAlertNotificationCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateAlertNotification(ctx context.Context, cmd *models.UpdateAlertNotificationCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateAlertNotificationWithUid(ctx context.Context, cmd *models.UpdateAlertNotificationWithUidCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SetAlertNotificationStateToCompleteCommand(ctx context.Context, cmd *models.SetAlertNotificationStateToCompleteCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) SetAlertNotificationStateToPendingCommand(ctx context.Context, cmd *models.SetAlertNotificationStateToPendingCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetOrCreateAlertNotificationState(ctx context.Context, cmd *models.GetOrCreateNotificationStateQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetAPIKeys(ctx context.Context, query *models.GetApiKeysQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetNonServiceAccountAPIKeys(ctx context.Context) []*models.ApiKey {
	return nil
}

func (m SQLStoreMock) DeleteApiKey(ctx context.Context, cmd *models.DeleteApiKeyCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) AddAPIKey(ctx context.Context, cmd *models.AddApiKeyCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateApikeyServiceAccount(ctx context.Context, apikeyId int64, saccountId int64) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetApiKeyById(ctx context.Context, query *models.GetApiKeyByIdQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetApiKeyByName(ctx context.Context, query *models.GetApiKeyByNameQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateTempUserStatus(ctx context.Context, cmd *models.UpdateTempUserStatusCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) CreateTempUser(ctx context.Context, cmd *models.CreateTempUserCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) UpdateTempUserWithEmailSent(ctx context.Context, cmd *models.UpdateTempUserWithEmailSentCommand) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetTempUsersQuery(ctx context.Context, query *models.GetTempUsersQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) GetTempUserByCode(ctx context.Context, query *models.GetTempUserByCodeQuery) error {
	return m.ExpectedError
}

func (m SQLStoreMock) ExpireOldUserInvites(ctx context.Context, cmd *models.ExpireTempUsersCommand) error {
	return m.ExpectedError
}
