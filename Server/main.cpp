#define CURL_STATICLIB
#include <iostream>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <string>
#include <stdio.h>
#include <curl/curl.h>


using namespace std;


int main(int argc, char** argv)
{

	string barcode = argv[1];
	string url_with_barcode = "http://world.openfoodfacts.org/api/v0/product/"+barcode+".json";
	string command = "curl "+url_with_barcode+" --output /Users/MattWilfert/Desktop/test.json";
	system(command.c_str());


	boost::property_tree::ptree pt;
	boost::property_tree::read_json("test.json", pt);

	boost::property_tree::basic_ptree<std::string,std::string>::const_iterator iter = pt.begin(),iterEnd = pt.end();

	string product_name = pt.get<string>("product.product_name");
	cout << product_name << endl;  

	int proteins = pt.get<int>("product.nutriments.proteins");  
	cout << "Proteins: " << proteins << endl;

	int carbohydrates = pt.get<int>("product.nutriments.carbohydrates_100g");  
	cout << "Carbohydates: " << carbohydrates << endl;

	int fats = pt.get<int>("product.nutriments.fat_100g");  
	cout << "Fats: " << fats << endl;

	//system("rm /Users/MattWilfert/Desktop/test.json");
}