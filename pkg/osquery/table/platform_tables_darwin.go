// +build darwin

package table

import (
	"github.com/go-kit/kit/log"
	"github.com/knightsc/system_policy/osquery/table/kextpolicy"
	"github.com/knightsc/system_policy/osquery/table/legacyexec"
	osquery "github.com/kolide/osquery-go"
	"github.com/kolide/osquery-go/plugin/table"
	_ "github.com/mattn/go-sqlite3"
)

func platformTables(client *osquery.ExtensionManagerClient, logger log.Logger) []*table.Plugin {
	munki := new(MunkiInfo)

	return []*table.Plugin{
		Airdrop(client),
		AppIcons(),
		BestPractices(client),
		ChromeLoginDataEmails(client, logger),
		ChromeLoginKeychainInfo(client, logger),
		ChromeUserProfiles(client, logger),
		EmailAddresses(client, logger),
		GDriveSyncConfig(client, logger),
		GDriveSyncHistoryInfo(client, logger),
		KolideVulnerabilities(client, logger),
		LauncherInfoTable(),
		MachoInfo(),
		MacOSUpdate(client),
		MDMInfo(logger),
		munki.ManagedInstalls(client, logger),
		munki.MunkiReport(client, logger),
		OnePasswordAccounts(client, logger),
		SlackConfig(client, logger),
		Spotlight(),
		UserAvatar(logger),
		legacyexec.TablePlugin(),
		kextpolicy.TablePlugin(),
	}
}
