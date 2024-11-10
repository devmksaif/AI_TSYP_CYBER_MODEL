
# AI-Powered Network Threat Detection API

This project integrates Artificial Intelligence into cybersecurity by monitoring and analyzing network traffic in real-time. It consists of two main components:

1. **Primary API (Golang-based)**: Captures and processes network packets sent or received by the server.
2. **AI Model API**: Analyzes the processed packets to classify and identify potential threats.

---

## 🚀 **Features**

- **Real-time Packet Capture**: Monitors all network traffic (sent and received packets).
- **User Identification**: Identifies the source of each packet (e.g., browser, user type - root, normal).
- **Batch Processing**: Processes 1000 packets at a time.
- **Threat Analysis**: Sends packet data to the AI Model API for threat classification.
- **Threat Notification**: Alerts when threat levels exceed 60%.
- **Threat Levels**:
  - **Green**: No threat detected.
  - **Yellow**: Low-level threat.
  - **Orange**: Medium-level threat.
  - **Red**: High-level threat.

---

## 🛠 **Technologies Used**

### Backend
- **Golang**: High-performance API for packet handling.
- **AI Model API**: Machine Learning-based threat detection.

### Tools
- **Packet Capture Library**: [Gopacket](https://github.com/google/gopacket) (or any packet processing library of your choice).
- **Machine Learning Framework**: TensorFlow, PyTorch, or other.

---

## 📂 Project Structure
```
├── .gitignore # Git ignore file ├── AI_Powered_Network_Threat_Detection_API.md # Project documentation ├── BinaryPrediction.ipynb # Jupyter notebook for binary prediction ├── capture.pcap # Packet capture file (Wireshark format) ├── go.mod # Go module file ├── go.sum # Go dependencies file ├── LICENSE # Project license ├── local-receive.json # Local JSON config for receiving data ├── local-send.json # Local JSON config for sending data ├── main.go # Entry point for the Primary API ├── Mdd.h5 # AI Model file ├── MulticlassPrediction.ipynb # Jupyter notebook for multiclass prediction ├── README.md # Main project README ├── Scaler.pkl # Scaler for preprocessing input data ├── wlan0-receive.json # JSON config for receiving data on wlan0 └── wlan0-send.json # JSON config for sending data on wlan0
```


## 🔧 **Setup and Installation**

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/devmksaif/AI_TSYP_CYBER_MODEL.git
   cd AI_TSYP_CYBER_MODEL
   ```

2. **Install Dependencies**  
   Ensure you have Go installed. Then run:
   ```bash
   go mod tidy
   ```

3. **Set Up Environment Variables**  
   Create a `.env` file and configure the following variables:
   ```env
   PORT=8080
   AI_API_URL=http://localhost:5000  # Replace with AI Model API URL
   ```

4. **Run the Application**  
   ```bash
   go run main.go
   ```

---

## 🧪 **Usage**

### 1. **Start the Primary API**  
   The API listens for incoming network traffic and processes 1000 packets at a time.

### 2. **Threat Detection Workflow**  
   - Captures and categorizes packets.
   - Sends packet data to the AI Model API.
   - Receives and logs threat levels.
   - Notifies you if the threat level exceeds 60%.

### 3. **Threat Levels Response Example**
```json
{
  "packet_id": 123,
  "source": "Browser",
  "user_type": "root",
  "threat_level": "orange",
  "threat_score": 72.5,
  "message": "Potential medium-level threat detected."
}
```

---

## 📝 **API Endpoints**

### **Primary API (Golang)**

- **`GET /packets`**  
  Returns the captured packets.

- **`POST /analyze`**  
  Sends the batch of packets to the AI Model API.

- **`GET /alerts`**  
  Returns a list of recent threat alerts.

### **AI Model API**

- **`POST /analyze-packets`**  
  Accepts a batch of packets and returns threat analysis.

---

## 🚀 **AI Model Integration**

The AI Model API uses a trained model to predict the threat level based on the processed network packets. The model is loaded from a trained file, performs batch predictions, and returns the threat level.

---

### 🛠 **AI Model Implementation Details**

1. **Loading the Model**:
   - The AI model is loaded from a `.h5` file (or another appropriate model format).
   - A scaler `.pkl` file is used to preprocess input data for consistent results.

2. **Packet Data Processing**:
   - Captured packets are preprocessed into a suitable format before being passed to the model.
   - Preprocessing includes feature extraction and scaling using the `.pkl` scaler file.

3. **Threat Classification**:
   - The model classifies packets into threat levels based on learned patterns.
   - Results are mapped to threat levels (Green, Yellow, Orange, Red) and returned to the Primary API.

---

### 📝 **AI Model API Endpoints**

- **`POST /analyze-packets`**
  - **Description**: Accepts a batch of packets and returns a JSON response with threat levels.
  - **Input**: JSON payload with packet data (up to 1000 packets).
  - **Output**: Threat classification for each packet in the batch.

#### Request Example
```json
{
  "packets": [
    {
      "protocol_type": "tcp",
      "service": "http",
      "flag": "SF",
      "src_bytes": 1032,
      "dst_bytes": 80
      // Additional required features
    },
    // More packet data
  ]
}
```

#### Response Example
```json
{
  "analysis": [
    {
      "packet_id": 1,
      "threat_level": "orange",
      "threat_score": 72.5,
      "message": "Threat level: medium with score 72.5%"
    }
    // More analysis results
  ]
}
```

---

### 🧩 **AI Model Code Example** (Python-based AI Model API)

Here’s a basic structure to demonstrate loading the model, preprocessing, and returning predictions:

```python
from flask import Flask, request, jsonify
import tensorflow as tf
import pickle
import numpy as np

app = Flask(__name__)

# Load the model and scaler
model = tf.keras.models.load_model('path/to/your_model.h5')
with open('path/to/scaler.pkl', 'rb') as f:
    scaler = pickle.load(f)

@app.route('/analyze-packets', methods=['POST'])
def analyze_packets():
    data = request.json["packets"]
    # Extract features for model input
    features = np.array([extract_features(packet) for packet in data])
    scaled_features = scaler.transform(features)
    
    # Predict threat levels
    predictions = model.predict(scaled_features)
    results = []

    for i, prediction in enumerate(predictions):
        threat_score = prediction * 100  # assuming prediction is between 0 and 1
        threat_level = classify_threat(threat_score)  # Green, Yellow, Orange, Red
        results.append({
            "packet_id": i,
            "threat_level": threat_level,
            "threat_score": threat_score,
            "message": f"Threat level: {threat_level} with score {threat_score:.2f}%"
        })

    return jsonify({"analysis": results})

def extract_features(packet):
    # Extract relevant features from packet
    # return in model's expected order as numpy array
    pass

def classify_threat(score):
    if score > 60:
        return "red"
    elif score > 40:
        return "orange"
    elif score > 20:
        return "yellow"
    return "green"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
```

---

### 🔧 **Running the AI Model API**

1. **Set up the Environment**:
   - Ensure all dependencies for TensorFlow and Flask are installed.
   - Place the model `.h5` and scaler `.pkl` files in the correct paths.

2. **Start the AI Model API**:
   ```bash
   python ai_model_api.py
   ```

This will start the AI Model API, listening on port 5000, ready to process and classify packets received from the Primary API.

---

### ⚙️ **Primary API and AI Model Integration**

In the Primary API (Golang), configure the `/analyze` endpoint to send packet batches to the AI Model API, parse the results, and handle threat notifications based on the AI model's analysis.
