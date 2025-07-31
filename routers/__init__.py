from fastapi import APIRouter

from .headers import header_router
from .boxart import box_art_router

api_router = APIRouter()

api_router.include_router(header_router, prefix="/header", tags=["header"])
api_router.include_router(box_art_router, prefix="/box-art", tags=["box-art"])