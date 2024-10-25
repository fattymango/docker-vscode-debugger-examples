import debugpy
from fastapi import FastAPI

# Start the debugger and listen on port 5678
debugpy.listen(("0.0.0.0", 5678))
print("Waiting for debugger to attach...")

# Uncomment this line to pause execution until the debugger attaches
# debugpy.wait_for_client()

app = FastAPI()


@app.get("/")
async def hello_world():
    return {"message": "Hello, World!"}
