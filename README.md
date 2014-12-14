ageing_files
============

Reports files that are beyond a certain age.

Downloads
---------
Download the [linux binary from S3](https://s3.amazonaws.com/emergeadapt-public-downloads/ageing_files).

Setup Crontab
-------------
Set up a crontab entry as follows:

    0,20,40 08-21 * * 1-6         root    ageing_files r --basepath /srv/ftp/viking --folders orders,transit,upload --maxage 1200 --exclude=FLGROUP.SEN | notify_hipchat --token 9ba63663685bc06d47f9a9ab137fb0 --room_id Infrastructure --from Axle --to ijonas,stewart --subject "Files older than 20 minutes"
    
The example above runs the checks

- every 20 minutes
- between 8am & 9pm
- Monday through to Saturday
- notifying Ijonas @ Stewart

