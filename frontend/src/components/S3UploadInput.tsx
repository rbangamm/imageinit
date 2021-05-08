import { FC, useContext, useEffect, useState } from 'react';
import S3 from 'aws-sdk/clients/s3';
import AWS from 'aws-sdk';
import { UserContext } from '../hooks/UserContext';

// Set the AWS Region
const REGION = "us-east-2"; //REGION

// Initialize the Amazon Cognito credentials provider
AWS.config.update({
  region: REGION,
  credentials: new AWS.CognitoIdentityCredentials({
    IdentityPoolId: String(process.env.REACT_APP_COGNITO_IDENTITY_ID), // IDENTITY_POOL_ID
  }),
});

const albumBucketName = "imageinitimages"; //BUCKET_NAME

var s3 = new AWS.S3({
    apiVersion: "2006-03-01",
    params: { Bucket: albumBucketName }
});

interface PhotoProps {
    photoUrl : string,
    albumName : string, 
    photoKey : string, 
    albumPhotosKey : string
}

const S3UploadInput = () => {

    const {user} = useContext(UserContext);

    const [selectedFile, setSelectedFile] = useState<File>();
    const [displayPhotos, setDisplayPhotos] = useState<PhotoProps[]>([]);
    const [message, setMessage] = useState("");

    const Photo : FC<PhotoProps> = ({photoUrl, 
        albumName, photoKey, albumPhotosKey}) => {
        return (
            <span style={{float: "left"}}>
                <div><img style={{width:128, height:128}} src={photoUrl} alt={photoUrl}/></div>
                <div>
                    <h3 onClick={(e) => {deletePhoto(albumName, photoKey)}}>X</h3>
                    <span>{photoKey.replace(albumPhotosKey, "")}</span>
                </div>
            </span>
        )
    }

    const viewAlbum = (albumName : string) => {
        var albumPhotosKey = encodeURIComponent(albumName) + "/";
        s3.listObjectsV2({ Prefix: user.username, Bucket: albumBucketName }, function(err, data) {
            if (err) {
            return alert("There was an error viewing your album: " + err.message);
          }

          let href = "https://s3." + REGION + ".amazonaws.com/";
          let bucketUrl = href + albumBucketName + "/";
      
            if (data && data.Contents) {
                let arr : PhotoProps[] = []
                data.Contents.forEach(photo => {
                    var photoKey = String(photo.Key);
                    var photoUrl = bucketUrl + encodeURIComponent(photoKey);
                    let p : PhotoProps = {photoUrl, albumName, photoKey, albumPhotosKey};
                    if (p !== null) {
                        arr.push(p);
                    }
                });
                setDisplayPhotos(arr);
                if (arr.length > 0) {
                    setMessage("Click 'X' under the image you would like to delete.");
                } else {
                    setMessage("You currently have no images.")
                }
            };
          });
      }

    const handleFileInput = (e : React.ChangeEvent<HTMLInputElement>) => {
        let file : File
        if (e) {
            if (e.target) {
                if (e.target.files) {
                    file = e.target.files[0]
                    setSelectedFile(file);
                }
            }
        }
    }

    const deletePhoto = (albumName : string, photoKey : string) => {
        s3.deleteObject({ Key: photoKey, Bucket : albumBucketName }, function(err, data) {
            if (err) {
            return alert("There was an error deleting your photo: " + err);
            }
            alert("Successfully deleted photo.");
            viewAlbum(albumName);
        });
    }

    useEffect(() => {
        if (selectedFile instanceof File) {
            var albumPhotosKey = encodeURIComponent(user.username) + "/";

            var photoKey = albumPhotosKey + selectedFile.name;

            // Use S3 ManagedUpload class as it supports multipart uploads
            var upload = new S3.ManagedUpload({
                params: {
                Bucket: albumBucketName,
                Key: photoKey,
                Body: selectedFile
                }
            });

            var promise = upload.promise();

            promise.then(
                function(data) {
                alert("Successfully uploaded photo.");
                console.log(data);
                viewAlbum(user.username);
                },
                function(err) {
                return alert("There was an error uploading your photo: " + err);
                }
            );
        } else {
            viewAlbum(user.username);
        }
    }, [selectedFile]);

    return (<div>
        <input type="file" accept="image/*" onChange={handleFileInput}/>
        <div id="photoDisplay">
            <h2>User: {user.username}</h2>
            <p>{message}</p>
            <div>
                {displayPhotos.map((p) => {return <Photo {...p}/>})}
            </div>
        </div>
    </div>)
}

export default S3UploadInput;