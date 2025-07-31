from fastapi import APIRouter

from .profilerouter import profile_router
from .gamerouter import game_router

api_router = APIRouter()

api_router.include_router(profile_router, prefix="/profile", tags=["profile"])
api_router.include_router(game_router, prefix="/game", tags=["game"])