certbot --nginx --email bogomolovs693@gmail.com --agree-tos --no-eff-email -d api.milutin-jewellery.com -d admin.milutin-jewellery.com -d milutin-jewellery.com
nginx -s reload
