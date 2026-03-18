# frontend-app/main.py
import os
from typing import Dict
from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn

class User(BaseModel):
    id: int
    name: str
    email: str

class AppConfig:
    def __init__(self, app_config: Dict):
        self.host = app_config.get('host', '0.0.0.0')
        self.port = app_config.get('port', 8000)

def create_app(app_config: Dict):
    app = FastAPI()
    
    # Define routes
    @app.post('/users/')
    async def create_user(user: User):
        return user
    
    @app.get('/users/{user_id}')
    async def get_user(user_id: int):
        # Simulate database query
        user = User(id=user_id, name='John Doe', email='john.doe@example.com')
        return user
    
    return app

if __name__ == '__main__':
    app_config = {
        'host': '0.0.0.0',
        'port': 8000
    }
    app = create_app(app_config)
    uvicorn.run(app, host=app_config['host'], port=app_config['port'])