public class Main {
	public static void maint(String args[]) {
		System.out.println("Hello");
	}

	public void test() throws RuntimeException {
		System.out.println("Test");
		new Thread(() -> {
			System.out.println("");
		});
	}

}
