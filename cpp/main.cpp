#include <exception>
#include <iostream>
#include "deephaven/client/highlevel/client.h"

using namespace deephaven::client::highlevel;

int main() {
  try {
	const char *server = "localhost:10000";

    auto client = Client::connect(server);

	// TODO: Check about information related to field syncing or ListFields.

    auto manager = client.getManager();

	/*
		Creation of handle1 omitted.
		The C++ client does not have a direct equivalent of importTable,
		and uses TableMaker or Arrow Flight instead.
	*/
	auto handle1 = TableHandle(/* omitted */);

    auto handle2 = manager.emptyTable(10);

	auto handle3 = handle2.where("a > 10 && b % 2 == 0");

	auto handle4 = handle1.merge("a", { handle3 }); // Note that merge is a method on the table handle here.

	auto combo_agg = aggCombo({
		aggSum("foo = a", "bar = c"),
		aggAvg("avgA = a", "avgB = b")
	});
	auto handle5 = handle4.by(combo_agg, { "g" });

	/*
		There is no equivalent of the snapshot method in the C++ client.
		The C++ client only supports reading table data through Arrow Flight.
		There is a convenience method to print the table data, though.
	*/
	std::cout << handle5.stream(true) << '\n';
  } catch (const std::exception &e) {
    std::cerr << "Caught exception: " << e.what() << '\n';
  }

  return 0;
}

