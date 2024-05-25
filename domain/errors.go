package domain

import "github.com/pkg/errors"

var (
	// ErrServerStart is an error that occurs when the server fails to start
	ErrServerStart = errors.New("failed to start server")
	// ErrServerStop is an error that occurs when the server fails to stop
	ErrServerStop = errors.New("failed to stop server")
	// ErrServerAccept is an error that occurs when the server fails to accept a connection
	ErrServerAccept = errors.New("failed to accept connection")
	// ErrServerRead is an error that occurs when the server fails to read from a connection
	ErrServerRead = errors.New("failed to read from connection")
	// ErrInvalidSetRequest is an error that occurs when the server receives an invalid SET request
	ErrInvalidGetRequest = errors.New("invalid GET request")
	// ErrErrorGettingKey is an error that occurs when the server fails to get a key
	ErrErrorGettingKey = errors.New("error getting key")
	// ErrUnableToProcessRequest is an error that occurs when the server is unable to process a request
	ErrUnableToProcessRequest = errors.New("unable to process request")
)
