#!/bin/bash

PROJECT="real-time-forum"

mkdir -p $PROJECT
mkdir -p $PROJECT/backend/{cmd/server,config,database,middleware,utils,router}

touch $PROJECT/backend/cmd/server/main.go
touch $PROJECT/backend/config/config.go
touch $PROJECT/backend/database/{database.go,schema.sql}
touch $PROJECT/backend/middleware/{auth.go,method.go,cors.go}
touch $PROJECT/backend/utils/{response.go,validator.go,password.go,session.go}
touch $PROJECT/backend/router/router.go
FEATURES=("auth" "posts" "comments" "reactions" "chat")

for feature in "${FEATURES[@]}"
do
    mkdir -p $PROJECT/backend/features/$feature

    touch $PROJECT/backend/features/$feature/{handler.go,service.go,repository.go,routes.go,dto.go,model.go}

    if [ "$feature" == "chat" ]; then
        touch $PROJECT/backend/features/chat/{websocket.go,client.go,hub.go}
    fi
done
mkdir -p $PROJECT/frontend/{css,assets}
mkdir -p $PROJECT/frontend/js/{shared,features}

touch $PROJECT/frontend/index.html
touch $PROJECT/frontend/js/{app.js,router.js}

touch $PROJECT/frontend/js/shared/{fetch.js,modal.js,helpers.js}
for feature in "${FEATURES[@]}"
do
    mkdir -p $PROJECT/frontend/js/features/$feature

    touch $PROJECT/frontend/js/features/$feature/{api.js,template.js,events.js}

    if [ "$feature" == "auth" ] || [ "$feature" == "posts" ]; then
        touch $PROJECT/frontend/js/features/$feature/page.js
    fi

    if [ "$feature" == "chat" ]; then
        touch $PROJECT/frontend/js/features/chat/websocket.js
    fi
done