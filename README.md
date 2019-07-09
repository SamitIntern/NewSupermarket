**Supermarket**
------------------




About the Project
------------------

The following project titled as **Supermarket** is a web application that is designed to manage the repository of the Supermarket.
Users can add, delete or view existing products in the repository using the application. Every product as record has certain associated characterestics with it that helps users identify the product. Every product has a unique product code and a unique ID along with a product name and price.
***

Using the Application
----------------------

In the following section, you will get to know how to use the application. Once you have cloned the repository, go to *$GOPATH/src/folder-name-you-cloned-it-to/* in your Goland terminal and type the following: 

`go run ./*.go`


*Starting the Application*

To start the application, curl the following URL:

`http://localhost:8080/supermarket/`

Your terminal should look like this: 


![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Welcome%20screen.png "Welcome Screen")

***


*Lets initialize the repository*

To initialize the repository, curl the following URL: 

`http://localhost:8080/supermarket/get/createrepo/`

Your terminal should look something like this:


![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Repository%20Creation.png "Repository Creation")


***


*Let's add a Product*

To add a new product to the repository, curl the following URL:

`http://localhost:8080/supermarket/new/`

Your terminal should look like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/valid%20add.png "Valid Addition")

Here *Samit* is the name of the product, that is priced at *$2.7* and has a product code of *ffff-gggg-dddd-vvv3*. The generated ID is 5 (this is generated within the code).

***

*Let's delete a Product*

To delete a product in the repository, curl following URL:

`http://localhost:8080/supermarket/delete/1/`

Your terminal should look like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Valid%20Delete.png "Valid Deletion")

Notice that the image where repository was created had a product with ID 2 but now it is no more there. 2 is the ID of the product that was to be deleted.

***

*Let's see the current state of our repository*

This is to see the current state of the repository i.e. what all products are in the repository currently. Curl the following URL:

`http://localhost:8080/supermarket/all/showall/`

Your terminal should look like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Repository%20Search.png "Repository Search")


***


*Let's look for a particular product in the repository*

To look for a particular product in the repository, curl the following URL:

`http://localhost:8080/supermarket/2/`

Your browser tab should look like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Product%20Search.png "Product Search")

Notice that the "2" in the end is the product ID.

***


Edge Cases
-----------

In the following section, let's explore a few edge cases where you might feel like you are using the application correctly but not getting the correct results. This also has solutions and reasons for the same.


*Deleting the same Product twice or more*

In case a product has already been deleted, if you try to delete it again, you may see something like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Invalid%20Delete.png "Invalid Deletion")

Reason: The product is no longer in the repository and hence the application cannot delete it.

***

*Adding the same product twice or more*

In case you try adding a product with the same *Product Code* again, you may see something like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Invalid%20Addition%201.png "Invalid Addition 1")

Reason: When you try adding a product twice, the application rejects it as a product with the same product code already exists in the repository. So you cannot add again. Also, products with the same name can exist but not product code. For example Cranberry Juice of volume 32 FL Oz and of volume 64 FL Oz will both be called Cranberry Juice but will have a different product code.

***

*Adding a Product with the same Product Code unknowingly*

There might be a scenario where you feel that your product code is unique and should be added, but that is not what's happening. There you might see something like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Invalid%20Addition%202.png "Invalid Addition 2")

Reason: The product codes are unique and case insensitive. Notice that the product code is the same in the previous image but the four Gs are in caps here. The application will treat these both codes as same and hance just by changing the case of your product code (upper to lower case or vice versa) you do not generate a new product code.

***


*Looking for a non-existent product in the repository*

Sometimes, you may search for a product in the repository and get an error. This generally would only happen if the repository no more contains that product, something like someone else using the application has already deleted it. In such a scenario, your browser screen should look something like this:

![alt text](https://github.com/SamitIntern/Supermarket2/blob/Samit/Screenshots/Invalid%20Search.png "Invalid Search")

Reason: Remember that while deleting a product we chose to delete the product with product ID as 1.

***


Conclusion
-----------

To summarize, this application caters to most of the needs of supermarket as requested. The application can be checked via  "Postman" for all the DELETE and POST methods. The GET services can be tested via Postman as well as via the browser.
