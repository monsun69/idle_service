#include <stdio.h>
#include "main.h"

__declspec(dllexport)
int main() {
}

__declspec(dllexport)
void ServiceMain() {
	EntryPoint();
}
