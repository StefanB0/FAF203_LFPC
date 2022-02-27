module com.lfpc.lfpc_code {
    requires javafx.controls;
    requires javafx.fxml;
    requires org.json;
    requires guru.nidi.graphviz;


    opens com.lfpc.lfpc_code to javafx.fxml;
    exports com.lfpc.lfpc_code;
}