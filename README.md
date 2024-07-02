# MOS3
`MOS3` is a mock application for Amazon S3, meaning `My Own S3`, pronounced `mɒsˈθri`.
# Install

```bash
docker run -p 3333:3333 -v ./up:/app/upload -it --rm mos3
```
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
