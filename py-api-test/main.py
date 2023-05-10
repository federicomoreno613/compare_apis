from fastapi import FastAPI
from fastapi import HTTPException
from fastapi import Request
from fastapi.responses import PlainTextResponse
from fastapi.responses import Response
from fastapi.encoders import jsonable_encoder
import time
import datetime
import uvicorn
import os

app = FastAPI()
counter = 0

@app.middleware("http")
async def increment_counter(request: Request, call_next):
    global counter
    counter += 1
    response = await call_next(request)
    return response

@app.post("/date", response_class=PlainTextResponse)
async def date_handler(full_format: bool = False):
    if full_format:
        date_str = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    else:
        date_str = datetime.datetime.now().strftime("%Y-%m-%d")
    return date_str

@app.get("/counter", response_class=PlainTextResponse)
async def counter_handler():
    return str(counter)

@app.exception_handler(HTTPException)
async def http_exception_handler(request, exc):
    return Response(content=exc.detail, status_code=exc.status_code)

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=int(os.environ.get("PORT", 8081)))
    