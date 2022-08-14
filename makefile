genmock:
	mockery --dir internal/project/ports --all --output internal/project/mocks --log-level="debug"
	mockery --dir internal/project/usecase --all --output internal/project/mocks --log-level="debug"
