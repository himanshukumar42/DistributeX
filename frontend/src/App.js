import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

const App = () => {
    const [file, setFile] = useState(null);
    const [fileList, setFileList] = useState([]);

    useEffect(() => {
        const fetchFiles = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/v1/files');
                if (response.data.status === 'success') {
                    setFileList(response.data.data || []); 
                }
            } catch (error) {
                console.error('Error fetching files:', error);
            }
        };

        fetchFiles();
    }, []);

    const handleFileChange = (e) => {
        setFile(e.target.files[0]);
    };

    const handleUpload = async () => {
        if (!file) return;
        const formData = new FormData();
        formData.append('file', file);

        try {
            await axios.post('http://localhost:8080/api/v1/upload', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });
            const response = await axios.get('http://localhost:8080/api/v1/files');
            if (response.data.status === 'success') {
                setFileList(response.data.data || []); 
            }
            setFile(null);
        } catch (error) {
            console.error('Error uploading file:', error);
        }
    };

    const handleDownload = (fileId) => {
        const downloadUrl = `http://localhost:8080/api/v1/download/${fileId}`;
        window.location.href = downloadUrl;
    };

    return (
        <div className="app">
            <h1>DistributeX - File Storage Server</h1>
            <div className="upload-container">
                <input 
                    type="file" 
                    onChange={handleFileChange} 
                    className="file-input" 
                />
                <button 
                    onClick={handleUpload} 
                    className="upload-button"
                >
                    Upload
                </button>
            </div>
            <div className="file-list">
                <h2>Uploaded Files</h2>
                {Array.isArray(fileList) && fileList.length === 0 ? (
                    <p>No files uploaded yet. Please upload a file.</p>
                ) : (
                    <table className="file-table">
                        <thead>
                            <tr>
                                <th>Filename</th>
                                <th>Parts</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {fileList.map((file) => (
                                <tr key={file.file_id}>
                                    <td>{file.filename}</td>
                                    <td>{file.part_count}</td>
                                    <td>
                                        <button 
                                            onClick={() => handleDownload(file.file_id)} 
                                            className="download-button"
                                        >
                                            Download
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}
            </div>
        </div>
    );
};

export default App;
