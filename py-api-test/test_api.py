import pytest
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def test_date_endpoint():
    response = client.post("/date?full_format=true")
    assert response.status_code == 200
    assert response.text != ""

    response = client.post("/date?full_format=false")
    assert response.status_code == 200
    assert response.text != ""

def test_counter_endpoint():
    response = client.get("/counter")
    assert response.status_code == 200
    assert response.text != ""
