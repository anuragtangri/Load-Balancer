from flask import Flask
from threading import Thread

app1 = Flask(__name__)
app2 = Flask(__name__)

# Define routes for each Flask app
@app1.route('/')
def hello_app1():
    return 'Hello from App 1!'

@app2.route('/')
def hello_app2():
    return 'Hello from App 2!'

# Run each Flask app in a separate thread
if __name__ == '__main__':
    thread1 = Thread(target=app1.run, kwargs={'port': 5000})
    thread2 = Thread(target=app2.run, kwargs={'port': 5001})

    thread1.start()
    thread2.start()
