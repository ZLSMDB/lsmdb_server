// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: etcd/v1/etcd.proto
#ifndef GRPC_etcd_2fv1_2fetcd_2eproto__INCLUDED
#define GRPC_etcd_2fv1_2fetcd_2eproto__INCLUDED

#include "etcd/v1/etcd.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace etcd {
namespace v1 {

class Etcd final {
 public:
  static constexpr char const* service_full_name() {
    return "etcd.v1.Etcd";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status Put(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::etcd::v1::EtcdPutReply* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdPutReply>> AsyncPut(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdPutReply>>(AsyncPutRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdPutReply>> PrepareAsyncPut(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdPutReply>>(PrepareAsyncPutRaw(context, request, cq));
    }
    virtual ::grpc::Status Get(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::etcd::v1::EtcdGetReply* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdGetReply>> AsyncGet(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdGetReply>>(AsyncGetRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdGetReply>> PrepareAsyncGet(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdGetReply>>(PrepareAsyncGetRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void Put(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest* request, ::etcd::v1::EtcdPutReply* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Put(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest* request, ::etcd::v1::EtcdPutReply* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      virtual void Get(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest* request, ::etcd::v1::EtcdGetReply* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Get(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest* request, ::etcd::v1::EtcdGetReply* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdPutReply>* AsyncPutRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdPutReply>* PrepareAsyncPutRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdGetReply>* AsyncGetRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::etcd::v1::EtcdGetReply>* PrepareAsyncGetRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status Put(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::etcd::v1::EtcdPutReply* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdPutReply>> AsyncPut(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdPutReply>>(AsyncPutRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdPutReply>> PrepareAsyncPut(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdPutReply>>(PrepareAsyncPutRaw(context, request, cq));
    }
    ::grpc::Status Get(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::etcd::v1::EtcdGetReply* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdGetReply>> AsyncGet(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdGetReply>>(AsyncGetRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdGetReply>> PrepareAsyncGet(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdGetReply>>(PrepareAsyncGetRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void Put(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest* request, ::etcd::v1::EtcdPutReply* response, std::function<void(::grpc::Status)>) override;
      void Put(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest* request, ::etcd::v1::EtcdPutReply* response, ::grpc::ClientUnaryReactor* reactor) override;
      void Get(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest* request, ::etcd::v1::EtcdGetReply* response, std::function<void(::grpc::Status)>) override;
      void Get(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest* request, ::etcd::v1::EtcdGetReply* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdPutReply>* AsyncPutRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdPutReply>* PrepareAsyncPutRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdPutRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdGetReply>* AsyncGetRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::etcd::v1::EtcdGetReply>* PrepareAsyncGetRaw(::grpc::ClientContext* context, const ::etcd::v1::EtcdGetRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Put_;
    const ::grpc::internal::RpcMethod rpcmethod_Get_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status Put(::grpc::ServerContext* context, const ::etcd::v1::EtcdPutRequest* request, ::etcd::v1::EtcdPutReply* response);
    virtual ::grpc::Status Get(::grpc::ServerContext* context, const ::etcd::v1::EtcdGetRequest* request, ::etcd::v1::EtcdGetReply* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Put : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Put() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Put() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Put(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestPut(::grpc::ServerContext* context, ::etcd::v1::EtcdPutRequest* request, ::grpc::ServerAsyncResponseWriter< ::etcd::v1::EtcdPutReply>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_Get : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Get() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_Get() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Get(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGet(::grpc::ServerContext* context, ::etcd::v1::EtcdGetRequest* request, ::grpc::ServerAsyncResponseWriter< ::etcd::v1::EtcdGetReply>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_Put<WithAsyncMethod_Get<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_Put : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Put() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::etcd::v1::EtcdPutRequest, ::etcd::v1::EtcdPutReply>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::etcd::v1::EtcdPutRequest* request, ::etcd::v1::EtcdPutReply* response) { return this->Put(context, request, response); }));}
    void SetMessageAllocatorFor_Put(
        ::grpc::MessageAllocator< ::etcd::v1::EtcdPutRequest, ::etcd::v1::EtcdPutReply>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::etcd::v1::EtcdPutRequest, ::etcd::v1::EtcdPutReply>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Put() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Put(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Put(
      ::grpc::CallbackServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_Get : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Get() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::etcd::v1::EtcdGetRequest, ::etcd::v1::EtcdGetReply>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::etcd::v1::EtcdGetRequest* request, ::etcd::v1::EtcdGetReply* response) { return this->Get(context, request, response); }));}
    void SetMessageAllocatorFor_Get(
        ::grpc::MessageAllocator< ::etcd::v1::EtcdGetRequest, ::etcd::v1::EtcdGetReply>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::etcd::v1::EtcdGetRequest, ::etcd::v1::EtcdGetReply>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Get() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Get(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Get(
      ::grpc::CallbackServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_Put<WithCallbackMethod_Get<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_Put : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Put() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Put() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Put(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_Get : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Get() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_Get() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Get(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_Put : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Put() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_Put() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Put(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestPut(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_Get : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Get() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_Get() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Get(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGet(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Put : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Put() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Put(context, request, response); }));
    }
    ~WithRawCallbackMethod_Put() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Put(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Put(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Get : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Get() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Get(context, request, response); }));
    }
    ~WithRawCallbackMethod_Get() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Get(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Get(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Put : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Put() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::etcd::v1::EtcdPutRequest, ::etcd::v1::EtcdPutReply>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::etcd::v1::EtcdPutRequest, ::etcd::v1::EtcdPutReply>* streamer) {
                       return this->StreamedPut(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Put() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Put(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdPutRequest* /*request*/, ::etcd::v1::EtcdPutReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedPut(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::etcd::v1::EtcdPutRequest,::etcd::v1::EtcdPutReply>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Get : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Get() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::etcd::v1::EtcdGetRequest, ::etcd::v1::EtcdGetReply>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::etcd::v1::EtcdGetRequest, ::etcd::v1::EtcdGetReply>* streamer) {
                       return this->StreamedGet(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Get() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Get(::grpc::ServerContext* /*context*/, const ::etcd::v1::EtcdGetRequest* /*request*/, ::etcd::v1::EtcdGetReply* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGet(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::etcd::v1::EtcdGetRequest,::etcd::v1::EtcdGetReply>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_Put<WithStreamedUnaryMethod_Get<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_Put<WithStreamedUnaryMethod_Get<Service > > StreamedService;
};

}  // namespace v1
}  // namespace etcd


#endif  // GRPC_etcd_2fv1_2fetcd_2eproto__INCLUDED
