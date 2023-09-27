FROM python:3.9-alpine

RUN python -m pip install poetry==1.6.1
RUN poetry config virtualenvs.in-project true

WORKDIR /app

# Install the dependencies
COPY pyproject.toml poetry.lock /app/
RUN poetry install --no-cache --no-root

# Install the package
COPY . .
RUN poetry install --no-cache --only-root

EXPOSE 9090
CMD ["poetry", "run", "python", "-m", "auth_service.main"]