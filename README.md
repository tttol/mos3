# MOS3
`MOS3` is a mock application for Amazon S3, meaning `My Own S3`, pronounced `mɒsˈθri`.
# Install
## Docker
Run command.
```bash
# https://hub.docker.com/r/tttol/mos3
docker run -p 3333:3333 -v ./upload:/app/upload -it --rm tttol:mos3
```
Then acccess http://localhost:3333/s3 .

## Docker Compose
Here is compose.yml.
```yml:compose.yml
version: '3.8'

services:
  mos3:
    image: tttol:mos3
    ports:
      - "3333:3333"
    volumes:
      - ./upload:/app/upload
```

Run command.
```bash
docker compose up -d
```

Then acccess http://localhost:3333/s3 .
# Usage
TBD
# TBD
- Accept requests from AWS SDK for Java (Priority: High)
  - **DL** 
    - [x] listObjectsV2
    - [x] getObject
  - **Upload**
    - [x] putObject
  - cp
    - [x] copyObject
  - rm
    - [x] deleteObject
- Accept requests from CLI (Priority: Low)
  - [x] ls
  - [x] cp
  - [ ] rm
  - [ ] mv
- Web GUI
  - TailwindCSS?
