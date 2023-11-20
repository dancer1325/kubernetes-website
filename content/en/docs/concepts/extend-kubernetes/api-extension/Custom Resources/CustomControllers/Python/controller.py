# myapp-controller.py
import kopf
from kubernetes import client, config

@kopf.on.create('example.com', 'v1', 'myapps')
@kopf.on.update('example.com', 'v1', 'myapps')
def process_myapp(body, **kwargs):
    # Print a message when a MyApp resource is created or updated
    print(f"MyApp resource {body['metadata']['name']} created/updated with spec: {body.get('spec', {})}")

@kopf.on.delete('example.com', 'v1', 'myapps')
def cleanup_myapp(name, **kwargs):
    # Print a message when a MyApp resource is deleted
    print(f"MyApp resource {name} deleted. Performing cleanup.")

if __name__ == "__main__":
    # Load the Kubernetes configuration from the default location
    config.load_kube_config()

    # Run the custom controller using kopf
    kopf.run(process_myapp)
