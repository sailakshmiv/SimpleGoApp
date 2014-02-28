SimpleGoApp
===========

This is a simple web application written in Go and a distelli manifest file that can be used to deploy your application to any server using the <a href="http://www.distelli.com" target="_blank">Distelli Deployment Platform</a>

PreRequisites:
---------------

* <a href="http://www.distelli.com/docs/setup.html" target="_blank">Install the DistelliCLI</a>
* <a href="http://www.distelli.com/docs/server-setup.html" target="_blank">Add your Server</a>. Make sure its an Ubuntu Server.

Deploying the SimpleNodeApp
---------------------------

Clone this repo and follow the steps below to deploy the SimpleGoApp to your server.

1. Edit the distelli-manifest.yml file and replace &lt;username&gt; at the top of the file with your username from your Distelli Account
2. Create the Application: (<a href="http://www.distelli.com/docs/creating-an-application.html" target="_blank"><i>What is this?</i></a>)

    <i>Replace &lt;username&gt; with your username</i>
    <pre>% distelli create &lt;username&gt;/SimpleGoApp
        Creating App: SimpleGoApp
        App SimpleGoApp created successfully
    </pre>

3. Build the application using the Go tools

    <pre>% go build -o bin/server src/server.go
    </pre>

4. Push a new Release: (<a href="http://www.distelli.com/docs/pushing-an-artifact.html" target="_blank"><i>what is a release?</i></a>)

    <pre>% distelli push -m "Initial Release of SimpleGoApp"
        Pushing App: SimpleGoApp
        Uploading: [==========] 100%
        Creating Releases... DONE
    </pre>

4. Deploy the release you just pushed using the Distelli Web Interface. (<a href="http://www.distelli.com/docs/starting-a-deployment.html" target="_blank"><i>show me how</i></a>)

Wait for the deployment to complete. Once the deployment is finished the SimpleGoApp will be running on your server on port <b>8080</b>. To test it out make a request and you should see a hello world message:

<pre>
curl http://ec2-67-202-47-54.compute-1.amazonaws.com:8080/
Hello World from Distelli!
</pre>
