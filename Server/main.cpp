#define CURL_STATICLIB
#include <iostream>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <string>
#include <stdio.h>
#include <curl/curl.h>
#include <cstdio>
#include <memory>


using namespace std;

int main(int argc, char** argv)
{


	while(1){

  		char buf[256],syscmd[512];
        int i;

        /* Get next barcode */
        printf("Waiting for bar code [q=quit]:  ");
        if (fgets(buf,255,stdin)==NULL)
            break;

        /* Clean CR/LF off of string */
        for (i=0;buf[i]!='\0' && buf[i]!='\r' && buf[i]!='\n';i++);
        buf[i]='\0';

        /* q = quit */
        if (!strcmp(buf,"q"))
            break;

		string url_with_barcode = "http://world.openfoodfacts.org/api/v0/product/"+std::string(buf)+".json";
		cout << url_with_barcode << endl;
		string command = "curl "+url_with_barcode+" --output /Users/MattWilfert/Desktop"+std::string(buf)+".json";
		system(command.c_str());

		 FILE *fp;
		  char measuredWeight[1035];

		  /* Open the command for reading. */
		  fp = popen("./usbscale/usbscale", "r");
		  if (fp == NULL) {
		    printf("Failed to run command\n" );
		    exit(1);
		  }

		  /* Read the output a line at a time - output it. */
		  while (fgets(measuredWeight, sizeof(measuredWeight)-1, fp) != NULL) {
		    printf("%s", measuredWeight);
		  }

		  /* close */
		  pclose(fp);

	


		boost::property_tree::ptree pt;
		boost::property_tree::read_json("/Users/MattWilfert/Desktop"+std::string(buf)+".json", pt);

		boost::property_tree::basic_ptree<std::string,std::string>::const_iterator iter = pt.begin(),iterEnd = pt.end();

		string serving = pt.get<string>("product.serving_size");
		int index_of_space = serving.find(" ");
		string temp = serving.substr(0, index_of_space);
		//int serving_size = stoi(temp);
		int serving_size = atoi(temp.c_str());
		string serving_type = serving.substr(index_of_space);

		int measuredWeightInt = atoi(measuredWeight);
		double ratio = (int)measuredWeightInt / (int)serving_size; 

		cout << "Serving Size: " << serving_size <<  serving_type << endl;
		cout << "Measured Weight: " << measuredWeight << endl;

		string product_name = pt.get<string>("product.product_name");
		cout << product_name << endl;  

		int proteins = pt.get<int>("product.nutriments.proteins");  
		cout << "Proteins: " << proteins * ratio << endl;

		int carbohydrates = pt.get<int>("product.nutriments.carbohydrates_100g");  
		cout << "Carbohydates: " << carbohydrates * ratio << endl;

		int fats = pt.get<int>("product.nutriments.fat_100g");  
		cout << "Fats: " << fats * ratio << endl;

		int calories = ratio*((fats * 9) + (proteins * 4) + (carbohydrates * 4));
		cout << "Calories: " << calories << endl;

		system(("rm /Users/MattWilfert/Desktop"+std::string(buf)+".json").c_str());
		replace(product_name.begin(), product_name.end(), ' ', '_');
		cout << product_name << endl;

		system(("python html_scraper.py "+std::string(product_name)+" "+std::to_string(measuredWeightInt)).c_str());
	}
}