package CassaSample

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	 "github.com/gocql/gocql"
	
	
)
	// THIS IS ADDED
	// log is the default package logger which we'll use to log
	var log = logger.GetLogger("activity-CassaSample")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
	// Get the activity data from the context
		clusterIP := context.GetInput("ClusterIP").(string)
		keySpace := context.GetInput("Keyspace").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo engine says [%s] to [%s]",clusterIP,keySpace)

	 // Provide the cassandra cluster instance here.
    cluster := gocql.NewCluster(clusterIP) 
 
    // gocql requires the keyspace to be provided before the session is created.
    // In future there might be provisions to do this later.
    cluster.Keyspace = keySpace
 
    // This is time after which the creation of session call would timeout.
    // This can be customised as needed.
    //cluster.Timeout = 10 * time.Second 
 
   // cluster.ProtoVersion = 4
    session, err := cluster.CreateSession()
	log.Debugf("Session Created Sucessfully")
	
	if err != nil {
       log.Debugf("Could not connect to cassandra cluster: ", err)
   }
	log.Debugf("Session : " , session)
	log.Debugf("Cluster : " , clusterIP)
	log.Debugf("Keyspace : ", keySpace)
	log.Debugf("Session Timeout : " ,cluster.Timeout)


	// Set the result as part of the context
	context.SetOutput("result", "The Flogo engine says Connection Successfull")

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
