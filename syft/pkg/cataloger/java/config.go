package java

type Config struct {
	SearchUnindexedArchives bool
	SearchIndexedArchives   bool
	SearchMavenForLicenses  bool
	MavenBaseURL            string
}
