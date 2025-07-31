from fastapi import APIRouter

from .profilerouter import profile_router

api_router = APIRouter()

api_router.include_router(profile_router, prefix="/profile", tags=["profile"])