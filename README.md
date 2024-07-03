# MOS3 - My Own S3
![GitHub](https://img.shields.io/github/license/tttol/mos3) ![GitHub](https://img.shields.io/github/v/release/tttol/mos3)   
`MOS3` is a mock application for Amazon S3, meaning `My Own S3`, pronounced `mɒsˈθri`.
# Install
## Docker
Run command.
```bash
# https://hub.docker.com/r/tttol/mos3
docker run -p 3333:3333 -v ./upload:/app/upload -it --rm tttol/mos3:latest
```
Then acccess http://localhost:3333/s3 .

## Docker Compose
compose.yml is [here](https://github.com/tttol/mos3/blob/main/compose.yml).  
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
