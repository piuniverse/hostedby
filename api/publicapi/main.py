from fastapi import FastAPI
import uvicorn
from motor.motor_asyncio import AsyncIOMotorClient
from config import Api, Database

from .routers import ip_router 


api = FastAPI()

@api.on_event("startup")
async def startup_db_client():
    api.mongodb_client = AsyncIOMotorClient(Database.DB_URL)
    api.mongodb = api.mongodb_client[Database.DB_NAME]

@api.on_event("shutdown")
async def shutdown_db_client():
    api.mongodb_client.close()

api.include_router(ip_router, tags=["ip"], prefix="/ip")

if __name__ == "__main__":
    uvicorn.run(
        "main:api",
        host=General.HOST,
        reload=General.DEBUG_MODE,
        port=General.PORT,
    )
