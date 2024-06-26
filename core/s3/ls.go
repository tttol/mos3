package s3

import (
	"fmt"
	"net/http"
	"time"
)

func Ls(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Date", time.Now().Format(time.RFC1123))
	w.Header().Set("Server", "MOS3")

	xmlResponse := `<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
    <Owner>
        <ID>1234567890123456789012345678901234567890123456789012345678901234</ID>
        <DisplayName>your-display-name</DisplayName>
    </Owner>
    <Buckets>
        <Bucket>
            <Name>example-bucket-1</Name>
            <CreationDate>2024-06-26T06:52:00.000Z</CreationDate>
        </Bucket>
        <Bucket>
            <Name>example-bucket-2</Name>
            <CreationDate>2024-06-25T14:20:00.000Z</CreationDate>
        </Bucket>
    </Buckets>
</ListAllMyBucketsResult>`

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(xmlResponse)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xmlResponse))
}
