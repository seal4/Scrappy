curl -H "User-agent: 'epic bot'" https://www.reddit.com/r/shitposting.json?limit=0 | jq '.' | grep url_overridden_by_dest | grep -Eo "https://.*()" > output.txt ;
#curl -H "User-agent: 'epic bot'" https://www.reddit.com/r/memes.json?limit=25 | jq '.' | grep url_overridden_by_dest | grep -Eo "https://.*()" >> output.txt ;
#curl -H "User-agent: 'epic bot'" https://www.reddit.com/r/okbuddyretard.json?limit=25 | jq '.' | grep url_overridden_by_dest | grep -Eo "https://.*()" >> output.txt ;
#curl -H "User-agent: 'epic bot'" https://www.reddit.com/r/dankmemes.json?limit=25 | jq '.' | grep url_overridden_by_dest | grep -Eo "https://.*()" >> output.txt ;
#curl -H "User-agent: 'epic bot'" https://www.reddit.com/r/okbuddyretard.json?limit=1 | jq '.' | grep url_overridden_by_dest | grep -Eo "https://.*()" > output.txt 

# Collects 100 memes