package main

import (
	"fmt"
	"log"
	"os"
)

type Job interface {
	ID() string
	Run() error
}

type BaseJob struct {
	id string
}

func (j *BaseJob) ID() string {
	return j.id
}

func (j *BaseJob) SetID(id string) {
	j.id = id
}

type CreateJob struct {
	BaseJob
	fileName string
	data     []byte
}

func NewCreateJob(id string, fileName string, data []byte) *CreateJob {
	return &CreateJob{
		BaseJob:  BaseJob{id},
		fileName: fileName,
		data:     data, // TODO: Copy data?
	}
}

func (j *CreateJob) Run() error {
	file, err := os.Create(j.fileName)
	if err != nil {
		return err
	}

	_, err = file.Write(j.data) // TODO: Check n == len(j.data)?
	file.Close()                // FIXME
	return err
}

type DeleteJob struct {
	BaseJob
	fileName string
}

func NewDeleteJob(id string, fileName string) *DeleteJob {
	return &DeleteJob{
		BaseJob:  BaseJob{id},
		fileName: fileName,
	}
}

func (j *DeleteJob) Run() error {
	return os.Remove(j.fileName)
}

func RunJobs(jobs []Job) error {
	for _, job := range jobs {
		fmt.Printf("running %s ", job.ID())
		err := job.Run()
		if err == nil {
			fmt.Println(" (OK)")
		} else {
			fmt.Printf(" (ERROR: %s)\n", err)
			return err
		}
	}

	return nil
}

func main() {
	jobs := []Job{
		NewCreateJob("j1", "/tmp/a", []byte("A data")),
		NewCreateJob("j2", "/tmp/b", []byte("B data")),
		NewDeleteJob("j3", "/tmp/b"),
	}
	if err := RunJobs(jobs); err != nil {
		log.Fatalf("error: %s", err)
	}
}
