# ageing_files
Reports files that are beyond a certain age.

## Downloads
Download the [linux binary from S3](https://s3.amazonaws.com/emergeadapt-public-downloads/ageing_files).

## Setup Crontab
Set up a crontab entry as follows:

    0,20,40 08-21 * * 1-6         root    ageing_files r --basepath /srv/ftp/home --folders orders,transit,upload --maxage 1200 --exclude=FLGROUP.SEN | notify_hipchat --token 9ba63663685bc06d47f9a9ab137fb0 --room_id Infrastructure --from Axle --to ijonas,stewart --subject "Files older than 20 minutes"
    
The example above runs the checks

- every 20 minutes
- between 8am & 9pm
- Monday through to Saturday
- notifying Ijonas @ Stewart

## Ageing_files Options
Ageing_files can be told to ignore files by using either exact file names or simple wildcard prefixes. Matching is case-insensitive.

     --exclude=filename2,wildcard*,filename2,wildcard2*

Folders are scanned using the *basepath* and *folders* parameters

     --basepath /srv/ftp/home --folders subfolder1,subfolder2,subfolder3

The maximum age of files is expressed in secods

     --maxage 1200

