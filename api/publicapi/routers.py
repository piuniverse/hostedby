import logging

from fastapi import APIRouter, Body, Request, HTTPException, status
from fastapi.responses import JSONResponse
from fastapi.encoders import jsonable_encoder
from pymongo.errors import DuplicateKeyError

from models.models import IPModel

logging.basicConfig(level=logging.INFO)

#IP address API
ip_router = APIRouter()


"""Get all IP Addresses"""
@ip_router.get("/", response_description="Get all IPs")
async def get_ips(request: Request):

    ips = []
    for doc in await request.app.mongodb["col_ips"].find().to_list(length=100):
        ips.append(doc)
    return ips

"""Get IP Address"""
@ip_router.get("/{id}", response_description="Get a single IP")
async def show_ip(id: str, request: Request):
    if (location := await request.app.mongodb["col_ips"].find_one({"_id": id})) is not None:
        return location

    raise HTTPException(status_code=404, detail=f"IP {id} not found")

