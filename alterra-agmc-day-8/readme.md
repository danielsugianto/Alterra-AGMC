DOCS: https://devcenter.heroku.com/articles/container-registry-and-runtime#pushing-an-existing-image
Step Build App:
1. go to heroku.com -> login -> Create new App
2. type app name and create app. -> alterra-agmc-ds
Step Deployment to Heroku:
1. Pull Docker Image -> pull docker image.png
2. Run docker tag <image> registry.heroku.com/<app>/<process-type> -> docker tag 9c5797818984 registry.heroku.com/alterra-agmc-ds/web
3. Run docker push registry.heroku.com/<app>/<process-type> -> docker push registry.heroku.com/alterra-agmc-ds/web
4. Run heroku container:release web -> heroku container: release web -a alterra-agmc-ds
