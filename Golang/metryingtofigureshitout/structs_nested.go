package main

type Repository struct {
	URL                    string
	Name                   string
	FilesystemPath         string
	LastScanned            string
	LinksScanned           int
	Links404               int
	HTMLReportFilename     string
	JSONReportFilename     string
	MetadataReportFileName string
	MarkdownFiles []struct {
		FileName      string
		FilePath      string
		HTTPAddr      string
		MarkdownLinks []struct {
			FileName      string
			LocalFilePath string
			HTTPFilePath  string
			Name          string
			Destination   string
			Type          string
			Status        string
		}
	}
}

func main (){

	r := Repository{}

	r1 := Repository{
		MarkdownFiles: []struct {
			FileName      string
			FilePath      string
			HTTPAddr      string
			MarkdownLinks []struct {
				FileName      string
				LocalFilePath string
				HTTPFilePath  string
				Name          string
				Destination   string
				Type          string
				Status        string
			}
		}{
			{
				"asd"
				{
					"fie"
				}
			},

		},
	}

	r.MarkdownFiles = r1.MarkdownFiles


}